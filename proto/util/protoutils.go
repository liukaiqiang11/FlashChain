package util

import (
	"crypto/sha256"
	"encoding/hex"

	pb "fchain/proto"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// CreateChaincodeProposal 创建链码请求
func CreateChaincodeProposal(cis *pb.ChaincodeSpec, creator []byte) (*pb.Proposal, string, error) {
	//创建一个随机的nonce
	nonce, err := getRandomNonce()
	if err != nil {
		return nil, "", err
	}

	// 计算txid
	txid := ComputeTxID(nonce, creator)

	return CreateChaincodeProposalWithTxIDNonce(txid, cis, nonce, creator)
}

// CreateChaincodeProposalWithTxID 创建链码请求（需要交易id）
func CreateChaincodeProposalWithTxID(cis *pb.ChaincodeSpec, creator []byte, txid string) (*pb.Proposal, string, error) {
	nonce, err := getRandomNonce()
	if err != nil {
		return nil, "", err
	}

	// 如果交易id为空，则计算交易id
	if txid == "" {
		txid = ComputeTxID(nonce, creator)
	}

	return CreateChaincodeProposalWithTxIDNonce(txid, cis, nonce, creator)
}

// CreateChaincodeProposalWithTxIDNonce 创建链码请求（需要交易id和nonce）
func CreateChaincodeProposalWithTxIDNonce(txid string, cis *pb.ChaincodeSpec, nonce, creator []byte) (*pb.Proposal, string, error) {

	cisBytes, err := proto.Marshal(cis)
	if err != nil {
		return nil, "", errors.Wrap(err, "error marshaling ChaincodeInvocationSpec")
	}

	hdr := &pb.SignatureHeader{
		Nonce:   nonce,
		Creator: creator,
	}

	hdrBytes, err := proto.Marshal(hdr)
	if err != nil {
		return nil, "", err
	}

	prop := &pb.Proposal{
		Header:  hdrBytes,
		Payload: cisBytes,
	}
	return prop, txid, nil
}

func CreateSignatureHeader(signer Signer) (*pb.SignatureHeader, error) {
	nonce, err := getRandomNonce()
	if err != nil {
		return nil, err
	}

	creator, err := signer.Serialize()
	if err != nil {
		return nil, err
	}

	hdr := &pb.SignatureHeader{
		Nonce:   nonce,
		Creator: creator,
	}
	return hdr, nil
}

func GetBytesProposalResponsePayload(hash []byte, response *pb.Response, result []byte) ([]byte, error) {
	cAct := &pb.ChaincodeAction{
		Results:  result,
		Response: response,
	}
	cActBytes, err := proto.Marshal(cAct)
	if err != nil {
		return nil, errors.Wrap(err, "error marshaling ChaincodeAction")
	}

	prp := &pb.ProposalResponsePayload{
		Extension:    cActBytes,
		ProposalHash: hash,
	}
	prpBytes, err := proto.Marshal(prp)
	return prpBytes, errors.Wrap(err, "error marshaling ProposalResponsePayload")
}

// GetBytesChaincodeSpec 获取 ChaincodeSpec 结构的字节数组
func GetBytesChaincodeSpec(cs *pb.ChaincodeSpec) ([]byte, error) {
	cppBytes, err := proto.Marshal(cs)
	return cppBytes, errors.Wrap(err, "error marshaling ChaincodeSpec")
}

// GetBytesChaincodeActionPayload 获取 ChaincodeActionPayload 结构的字节数组
func GetBytesChaincodeActionPayload(cap *pb.ChaincodeActionPayload) ([]byte, error) {
	cppBytes, err := proto.Marshal(cap)
	return cppBytes, errors.Wrap(err, "error marshaling ChaincodeActionPayload")
}

// GetBytesResponse 获取 Response 结构的字节数组
func GetBytesResponse(res *pb.Response) ([]byte, error) {
	resBytes, err := proto.Marshal(res)
	return resBytes, errors.Wrap(err, "error marshaling Response")
}

// GetBytesProposalResponse 获取 ProposalResponse 结构的字节数组
func GetBytesProposalResponse(pr *pb.ProposalResponse) ([]byte, error) {
	respBytes, err := proto.Marshal(pr)
	return respBytes, errors.Wrap(err, "error marshaling ProposalResponse")
}

// GetBytesSignatureHeader 获取 SignatureHeader 结构的字节数组
func GetBytesSignatureHeader(shdr *pb.SignatureHeader) ([]byte, error) {
	bytes, err := proto.Marshal(shdr)
	return bytes, errors.Wrap(err, "error marshaling SignatureHeader")
}

// GetBytesTransaction 获取 Transaction 结构的字节数组
func GetBytesTransaction(tx *pb.Transaction) ([]byte, error) {
	bytes, err := proto.Marshal(tx)
	return bytes, errors.Wrap(err, "error unmarshalling Transaction")
}

// GetBytesPayload 获取 Payload 结构的字节数组
func GetBytesPayload(payl *pb.Payload) ([]byte, error) {
	bytes, err := proto.Marshal(payl)
	return bytes, errors.Wrap(err, "error marshaling Payload")
}

// GetBytesEnvelope 获取 Envelope 结构的字节数组
func GetBytesEnvelope(env *pb.Envelope) ([]byte, error) {
	bytes, err := proto.Marshal(env)
	return bytes, errors.Wrap(err, "error marshaling Envelope")
}

func GetActionPayloadFromEnvelope(envBytes []byte) (*pb.ChaincodeActionPayload, error) {
	env, err := UnmarshalEnvelope(envBytes)
	if err != nil {
		return nil, err
	}
	return GetActionPayloadFromEnvelopeMsg(env)
}

func GetActionPayloadFromEnvelopeMsg(env *pb.Envelope) (*pb.ChaincodeActionPayload, error) {
	payl, err := UnmarshalPayload(env.Payload)
	if err != nil {
		return nil, err
	}

	tx, err := UnmarshalTransaction(payl.Data)
	if err != nil {
		return nil, err
	}

	if tx.Payload == nil {
		return nil, errors.New("at least one TransactionAction required")
	}

	respPayload := tx.Payload
	return respPayload, err
}

// CheckTxID 检测交易id是否正确
func CheckTxID(txid string, nonce, creator []byte) error {
	computedTxID := ComputeTxID(nonce, creator)

	if txid != computedTxID {
		return errors.Errorf("invalid txid. got [%s], expected [%s]", txid, computedTxID)
	}

	return nil
}

// ComputeTxID 计算交易id
func ComputeTxID(nonce, creator []byte) string {
	hasher := sha256.New()
	hasher.Write(nonce)
	hasher.Write(creator)
	return hex.EncodeToString(hasher.Sum(nil))
}

// CreateChaincodeTxID 创建交易id
func CreateChaincodeTxID(creator []byte) (string, error) {
	nonce, err := CreateNonce()
	if err != nil {
		return "", err
	}
	txid := ComputeTxID(nonce, creator)

	return txid, nil

}

func InvokedChaincodeName(proposalBytes []byte) (string, error) {
	proposal := &pb.Proposal{}
	err := proto.Unmarshal(proposalBytes, proposal)
	if err != nil {
		return "", errors.WithMessage(err, "could not unmarshal proposal")
	}

	cs := &pb.ChaincodeSpec{}
	err = proto.Unmarshal(proposal.Payload, cs)
	if err != nil {
		return "", errors.WithMessage(err, "could not unmarshal chaincode spec")
	}

	if cs.ChaincodeID == nil {
		return "", errors.Errorf("chaincode id is nil")
	}

	return cs.ChaincodeID.Name, nil
}
