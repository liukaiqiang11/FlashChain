package util

import (
	"bytes"
	"crypto/sha256"
	pb "fchain/proto"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GetChaincodeActionPayloadFromTransaction 从Transaction中获取ChaincodeActionPayload
func GetChaincodeActionPayloadFromTransaction(tx *pb.Transaction) (*pb.ChaincodeActionPayload, error) {
	if tx.Payload == nil {
		return nil, errors.New("no ChaincodeActionPayload in Transaction")
	}

	ccPayload := tx.Payload

	return ccPayload, nil
}

// GetProposalResponsePayloadFromTransaction 从Transaction中获取ProposalResponsePayload
func GetProposalResponsePayloadFromTransaction(tx *pb.Transaction) (*pb.ProposalResponsePayload, error) {

	ccPayload, err := GetChaincodeActionPayloadFromTransaction(tx)
	if err != nil {
		return nil, errors.New("no ChaincodeActionPayload in Transaction")
	}

	if ccPayload.ProposalResponsePayload == nil {
		return nil, errors.New("no ProposalResponsePayload in ChaincodeActionPayload")
	}

	prpBytes := ccPayload.ProposalResponsePayload
	prp, err := UnmarshalProposalResponsePayload(prpBytes)
	if err != nil {
		return nil, errors.New("Unmarshal ProposalResponsePayload error")
	}

	return prp, nil
}

// CreateSignedEnvelope 新建一个附带签名的Envelope
func CreateSignedEnvelope(
	signer Signer,
	dataMsg proto.Message,
) (*pb.Envelope, error) {
	var err error
	payloadSignatureHeader := &pb.SignatureHeader{}

	if signer != nil {
		payloadSignatureHeader, err = NewSignatureHeader(signer)
		if err != nil {
			return nil, err
		}
	}

	data, err := proto.Marshal(dataMsg)
	if err != nil {
		return nil, errors.Wrap(err, "error marshaling")
	}

	paylBytes := MarshalOrPanic(
		&pb.Payload{
			SignatureHeader: payloadSignatureHeader,
			Data:            data,
		},
	)

	var sig []byte
	if signer != nil {
		sig, err = signer.Sign(paylBytes)
		if err != nil {
			return nil, err
		}
	}

	env := &pb.Envelope{
		Payload:   paylBytes,
		Signature: sig,
	}

	return env, nil
}

type Signer interface {
	Sign(msg []byte) ([]byte, error)
	Serialize() ([]byte, error)
}

// CreateSignedTx 创建一笔附带签名的交易
func CreateSignedTx(
	txID string,
	proposal *pb.Proposal,
	signer Signer,
	resps ...*pb.ProposalResponse,
) (*pb.Envelope, error) {
	if len(resps) == 0 {
		return nil, errors.New("at least one proposal response is required")
	}

	if signer == nil {
		return nil, errors.New("signer is required when creating a signed transaction")
	}

	// the original header
	shdr, err := UnmarshalSignatureHeader(proposal.Header)
	if err != nil {
		return nil, err
	}

	// the original payload
	cs, err := UnmarshalChaincodeSpec(proposal.Payload)
	if err != nil {
		return nil, err
	}

	// check that the signer is the same that is referenced in the header
	signerBytes, err := signer.Serialize()
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(signerBytes, shdr.Creator) {
		return nil, errors.New("signer must be the same as the one referenced in the header")
	}

	// ensure that all actions are bitwise equal and that they are successful
	var a1 []byte
	for n, r := range resps {
		if r.Response.Status < 200 || r.Response.Status >= 400 {
			return nil, errors.Errorf("proposal response was not successful, error code %d, msg %s", r.Response.Status, r.Response.Message)
		}

		if n == 0 {
			a1 = r.Payload
			continue
		}

		if !bytes.Equal(a1, r.Payload) {
			return nil, errors.Errorf("Endorsement results are inconsistent")
		}
	}

	// fill endorsements according to their uniqueness
	endorsersUsed := make(map[string]struct{})
	var endorsements []*pb.Endorsement
	for _, r := range resps {
		if r.Endorsement == nil {
			continue
		}
		key := string(r.Endorsement.Endorser)
		if _, used := endorsersUsed[key]; used {
			continue
		}
		endorsements = append(endorsements, r.Endorsement)
		endorsersUsed[key] = struct{}{}
	}

	if len(endorsements) == 0 {
		return nil, errors.Errorf("no endorsements")
	}

	// create a transaction
	caPayload := &pb.ChaincodeActionPayload{Input: cs, Endorsements: endorsements, ProposalResponsePayload: resps[0].Payload}
	tx := &pb.Transaction{Payload: caPayload}

	// serialize the tx
	txBytes, err := GetBytesTransaction(tx)
	if err != nil {
		return nil, err
	}

	// create the payload
	payl := &pb.Payload{SignatureHeader: shdr, Data: txBytes}
	paylBytes, err := GetBytesPayload(payl)
	if err != nil {
		return nil, err
	}

	// sign the payload
	sig, err := signer.Sign(paylBytes)
	if err != nil {
		return nil, err
	}

	// here's the envelope
	return &pb.Envelope{TxID: txID, Payload: paylBytes, Signature: sig}, nil
}

// CreateProposalResponse 新建一个提案回复（ProposalResponse）
func CreateProposalResponse(
	shdrbytes []byte,
	payload []byte,
	response *pb.Response,
	results []byte,
	signingEndorser Signer,
) (*pb.ProposalResponse, error) {
	shdr, err := UnmarshalSignatureHeader(shdrbytes)
	if err != nil {
		return nil, err
	}

	pHashBytes, err := GetProposalHash(shdr, payload)
	if err != nil {
		return nil, errors.WithMessage(err, "error computing proposal hash")
	}

	// get the bytes of the proposal response payload - we need to sign them
	prpBytes, err := GetBytesProposalResponsePayload(pHashBytes, response, results)
	if err != nil {
		return nil, err
	}

	// serialize the signing IDentity
	endorser, err := signingEndorser.Serialize()
	if err != nil {
		return nil, errors.WithMessage(err, "error serializing signing IDentity")
	}

	// sign the concatenation of the proposal response and the serialized
	// endorser IDentity with this endorser's key
	signature, err := signingEndorser.Sign(append(prpBytes, endorser...))
	if err != nil {
		return nil, errors.WithMessage(err, "could not sign the proposal response payload")
	}

	resp := &pb.ProposalResponse{
		Version:   1,
		Timestamp: timestamppb.Now(),
		Endorsement: &pb.Endorsement{
			Signature: signature,
			Endorser:  endorser,
		},
		Payload: prpBytes,
		Response: &pb.Response{
			Status:  200,
			Message: "OK",
		},
	}

	return resp, nil
}

// CreateProposalResponseFailure 当背书提案失败或者链码失败时，创建一个 ProposalResponse
func CreateProposalResponseFailure(
	hdrbytes []byte,
	payl []byte,
	response *pb.Response,
	results []byte,
) (*pb.ProposalResponse, error) {
	shdr, err := UnmarshalSignatureHeader(hdrbytes)
	if err != nil {
		return nil, err
	}

	// obtain the proposal hash given proposal header, payload and the requested visibility
	pHashBytes, err := GetProposalHash(shdr, payl)
	if err != nil {
		return nil, errors.WithMessage(err, "error computing proposal hash")
	}

	// get the bytes of the proposal response payload
	prpBytes, err := GetBytesProposalResponsePayload(pHashBytes, response, results)
	if err != nil {
		return nil, err
	}

	resp := &pb.ProposalResponse{
		// Timestamp: TODO!
		Payload:  prpBytes,
		Response: response,
	}

	return resp, nil
}

// GetBytesProposalPayloadForTx 从交易中获取ProposalPayload的字节数组
func GetBytesProposalPayloadForTx(
	payload *pb.ChaincodeSpec,
) ([]byte, error) {
	// check for nil argument
	if payload == nil {
		return nil, errors.New("nil arguments")
	}

	// strip the transient bytes off the payload
	cppNoTransient := &pb.ChaincodeSpec{Input: payload.Input}
	cppBytes, err := GetBytesChaincodeSpec(cppNoTransient)
	if err != nil {
		return nil, err
	}

	return cppBytes, nil
}

// GetProposalHash 获取Proposal的hash
func GetProposalHash(signatureHeader *pb.SignatureHeader, ccPropPayl []byte) ([]byte, error) {
	// check for nil argument
	if signatureHeader == nil ||
		ccPropPayl == nil {
		return nil, errors.New("nil arguments")
	}

	// unmarshal the chaincode proposal payload
	cs, err := UnmarshalChaincodeSpec(ccPropPayl)
	if err != nil {
		return nil, err
	}

	ppBytes, err := GetBytesProposalPayloadForTx(cs)
	if err != nil {
		return nil, err
	}

	signatureHeaderData, err := proto.Marshal(signatureHeader)
	if err != nil {
		return nil, err
	}

	hash2 := sha256.New()
	// hash the serialized Signature Header object
	hash2.Write(signatureHeaderData)
	// hash of the part of the chaincode proposal payload that will go to the tx
	hash2.Write(ppBytes)
	return hash2.Sum(nil), nil
}

// GetSignedProposal 获取SignedProposal
func GetSignedProposal(proposal *pb.Proposal, signer Signer) (*pb.SignedProposal, error) {
	// check for nil argument
	if proposal == nil {
		return nil, errors.New("proposal cannot be nil")
	}

	if signer == nil {
		return nil, errors.New("signer cannot be nil")
	}

	proposalBytes, err := proto.Marshal(proposal)
	if err != nil {
		return nil, errors.Wrap(err, "error marshaling proposal")
	}

	signature, err := signer.Sign(proposalBytes)
	if err != nil {
		return nil, err
	}

	return &pb.SignedProposal{
		ProposalBytes: proposalBytes,
		Signature:     signature,
	}, nil
}

func GetEndorserMsgFromEnvelope(envelop *pb.Envelope) ([]byte, error) {
	tx, err := GetTransactionFromEnvelope(envelop)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting tx from envelop")
	}

	csBytes, err := GetBytesChaincodeSpec(tx.Payload.Input)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting chaincodeSpecBytes from tx")
	}
	payload, err := UnmarshalPayload(envelop.Payload)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting txPayload from envelop payload")
	}
	pHashBytes, err := GetProposalHash(payload.SignatureHeader, csBytes)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting proposalHash")
	}

	ca, err := GetChaincodeActionFromEnvelope(envelop)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting chaincodeAction from envelop")
	}
	prpBytes, err := GetBytesProposalResponsePayload(pHashBytes, ca.Response, ca.Results)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting txPayload from envelop")
	}
	return prpBytes, nil
}

func GetPayloadFromEnvelope(envelop *pb.Envelope) (*pb.Payload, error) {
	txPayload, err := UnmarshalPayload(envelop.Payload)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting txPayload from envelop")
	}
	return txPayload, nil
}

func GetTransactionFromEnvelope(envelop *pb.Envelope) (*pb.Transaction, error) {
	txPayload, err := GetPayloadFromEnvelope(envelop)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting txPayload from envelop")
	}
	if txPayload.Data == nil {
		return nil, errors.New("error getting txPayload data: payload data is nil")
	}

	tx, err := UnmarshalTransaction(txPayload.Data)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting transaction from envelop")
	}

	return tx, nil
}

func GetProposalResponsePayloadFromEnvelope(envelop *pb.Envelope) (*pb.ProposalResponsePayload, error) {
	tx, err := GetTransactionFromEnvelope(envelop)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting tx from envelop")
	}

	if tx.Payload == nil {
		return nil, errors.New("error getting tx Payload : tx Payload is nil")
	}

	proposalResponsePayload, err := UnmarshalProposalResponsePayload(tx.Payload.ProposalResponsePayload)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting proposalResponsePayload from envelop")
	}

	return proposalResponsePayload, nil
}

func GetChaincodeActionFromEnvelope(envelope *pb.Envelope) (*pb.ChaincodeAction, error) {
	proposalResponsePayload, err := GetProposalResponsePayloadFromEnvelope(envelope)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting proposalResponsePayload from envelop")
	}

	if proposalResponsePayload.Extension == nil {
		return nil, errors.New("error getting proposalResponsePayload Extension: proposalResponsePayload Extension is nil")
	}

	ca, err := UnmarshalChaincodeAction(proposalResponsePayload.Extension)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting chaincodeAction from envelop")
	}

	return ca, nil
}

func GetKVRWSetKwFromEnvelope(envelop *pb.Envelope) (*pb.KVRWSet, error) {

	ca, err := GetChaincodeActionFromEnvelope(envelop)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting ChaincodeAction from Envelope")
	}

	kVRWSet, err := UnmarshalKVRWSet(ca.Results)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting kVRWSet ChaincodeAction")
	}
	return kVRWSet, nil
}

func GetOrComputeTxIDFromEnvelope(envelop *pb.Envelope) (string, error) {

	if envelop.TxID != "" {
		return envelop.TxID, nil
	}

	txPayload, err := UnmarshalPayload(envelop.Payload)
	if err != nil {
		return "", errors.WithMessage(err, "error getting txID from payload")
	}

	if txPayload.SignatureHeader == nil {
		return "", errors.New("error getting txID from signatureHeader: payload signatureHeader is nil")
	}

	sighdr := txPayload.SignatureHeader

	txID := ComputeTxID(sighdr.Nonce, sighdr.Creator)
	return txID, nil
}

// GetOrComputeTxIDFromEnvelopeBytes 从Envelope中获取交易ID或者计算交易ID
func GetOrComputeTxIDFromEnvelopeBytes(txEnvelopBytes []byte) (string, error) {
	txEnvelope, err := UnmarshalEnvelope(txEnvelopBytes)
	if err != nil {
		return "", errors.WithMessage(err, "error getting txID from envelope")
	}

	return GetOrComputeTxIDFromEnvelope(txEnvelope)
}
