package util

import (
	pb "fchain/proto"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// UnmarshalBlock 将字节数组转为 Block 结构体
func UnmarshalBlock(encoded []byte) (*pb.Block, error) {
	block := &pb.Block{}
	err := proto.Unmarshal(encoded, block)
	return block, errors.Wrap(err, "error unmarshalling Block")
}

// UnmarshalChaincodeSpec 将字节数组转为 ChaincodeSpec 结构体
func UnmarshalChaincodeSpec(encoded []byte) (*pb.ChaincodeSpec, error) {
	cis := &pb.ChaincodeSpec{}
	err := proto.Unmarshal(encoded, cis)
	return cis, errors.Wrap(err, "error unmarshalling ChaincodeInvocationSpec")
}

// UnmarshalChaincodeInput 将字节数组转为 UnmarshalChaincodeInput 结构体
func UnmarshalChaincodeInput(encoded []byte) (*pb.ChaincodeInput, error) {
	ci := &pb.ChaincodeInput{}
	err := proto.Unmarshal(encoded, ci)
	return ci, errors.Wrap(err, "error unmarshalling ChaincodeInvocationSpec")
}

// UnmarshalPayload 将字节数组转为 Payload 结构体
func UnmarshalPayload(encoded []byte) (*pb.Payload, error) {
	payload := &pb.Payload{}
	err := proto.Unmarshal(encoded, payload)
	return payload, errors.Wrap(err, "error unmarshalling Payload")
}

// UnmarshalEnvelope 将字节数组转为 Envelope 结构体
func UnmarshalEnvelope(encoded []byte) (*pb.Envelope, error) {
	envelope := &pb.Envelope{}
	err := proto.Unmarshal(encoded, envelope)
	return envelope, errors.Wrap(err, "error unmarshalling Envelope")
}

// UnmarshalEnvelopes 将字节数组转为 Envelopes 结构体
func UnmarshalEnvelopes(encoded []byte) (*pb.Envelopes, error) {
	envelopes := &pb.Envelopes{}
	err := proto.Unmarshal(encoded, envelopes)
	return envelopes, errors.Wrap(err, "error unmarshalling Envelopes")
}

// UnmarshalChaincodeID 将字节数组转为 ChaincodeID 结构体
func UnmarshalChaincodeID(bytes []byte) (*pb.ChaincodeID, error) {
	ccid := &pb.ChaincodeID{}
	err := proto.Unmarshal(bytes, ccid)
	return ccid, errors.Wrap(err, "error unmarshalling ChaincodeID")
}

// UnmarshalSignatureHeader 将字节数组转为 SignatureHeader 结构体
func UnmarshalSignatureHeader(bytes []byte) (*pb.SignatureHeader, error) {
	sh := &pb.SignatureHeader{}
	err := proto.Unmarshal(bytes, sh)
	return sh, errors.Wrap(err, "error unmarshalling SignatureHeader")
}

// UnmarshalProposalResponse 将字节数组转为 ProposalResponse 结构体
func UnmarshalProposalResponse(prBytes []byte) (*pb.ProposalResponse, error) {
	proposalResponse := &pb.ProposalResponse{}
	err := proto.Unmarshal(prBytes, proposalResponse)
	return proposalResponse, errors.Wrap(err, "error unmarshalling ProposalResponse")
}

// UnmarshalChaincodeActionPayload 将字节数组转为 ChaincodeActionPayload 结构体
func UnmarshalChaincodeActionPayload(caBytes []byte) (*pb.ChaincodeActionPayload, error) {
	chaincodeActionPayload := &pb.ChaincodeActionPayload{}
	err := proto.Unmarshal(caBytes, chaincodeActionPayload)
	return chaincodeActionPayload, errors.Wrap(err, "error unmarshalling ChaincodeAction")
}

// UnmarshalResponse 将字节数组转为 Response 结构体
func UnmarshalResponse(resBytes []byte) (*pb.Response, error) {
	response := &pb.Response{}
	err := proto.Unmarshal(resBytes, response)
	return response, errors.Wrap(err, "error unmarshalling Response")
}

// UnmarshalProposalResponsePayload 将字节数组转为 ProposalResponsePayload 结构体
func UnmarshalProposalResponsePayload(prpBytes []byte) (*pb.ProposalResponsePayload, error) {
	prp := &pb.ProposalResponsePayload{}
	err := proto.Unmarshal(prpBytes, prp)
	return prp, errors.Wrap(err, "error unmarshalling ProposalResponsePayload")
}

// UnmarshalProposal 将字节数组转为 Proposal 结构体
func UnmarshalProposal(propBytes []byte) (*pb.Proposal, error) {
	prop := &pb.Proposal{}
	err := proto.Unmarshal(propBytes, prop)
	return prop, errors.Wrap(err, "error unmarshalling Proposal")
}

// UnmarshalTransaction 将字节数组转为 Transaction 结构体
func UnmarshalTransaction(txBytes []byte) (*pb.Transaction, error) {
	tx := &pb.Transaction{}
	err := proto.Unmarshal(txBytes, tx)
	return tx, errors.Wrap(err, "error unmarshalling Transaction")
}

// UnmarshalChaincodeAction 将字节数组转为 ChaincodeAction 结构体
func UnmarshalChaincodeAction(capBytes []byte) (*pb.ChaincodeAction, error) {
	ca := &pb.ChaincodeAction{}
	err := proto.Unmarshal(capBytes, ca)
	return ca, errors.Wrap(err, "error unmarshalling ChaincodeActionPayload")
}

// UnmarshalKVRWSet 将字节数组转为 KVRWSet 结构体
func UnmarshalKVRWSet(bytes []byte) (*pb.KVRWSet, error) {
	rws := &pb.KVRWSet{}
	err := proto.Unmarshal(bytes, rws)
	return rws, errors.Wrap(err, "error unmarshalling KVRWSet")
}

func UnmarshalCreator(bytes []byte) (*pb.Creator, error) {
	rws := &pb.Creator{}
	err := proto.Unmarshal(bytes, rws)
	return rws, errors.Wrap(err, "error unmarshalling Creator")
}

// UnmarshalEnvelopeOrPanic 将字节数组转为 Envelope 结构体，如果出现错误，则抛出panic
func UnmarshalEnvelopeOrPanic(encoded []byte) *pb.Envelope {
	envelope, err := UnmarshalEnvelope(encoded)
	if err != nil {
		panic(err)
	}
	return envelope
}

// UnmarshalBlockOrPanic 将字节数组转为 Block 结构体，如果出现错误，则抛出panic
func UnmarshalBlockOrPanic(encoded []byte) *pb.Block {
	block, err := UnmarshalBlock(encoded)
	if err != nil {
		panic(err)
	}
	return block
}

// UnmarshalSignatureHeaderOrPanic 将字节数组转为 SignatureHeader 结构体，如果出现错误，则抛出panic
func UnmarshalSignatureHeaderOrPanic(bytes []byte) *pb.SignatureHeader {
	sighdr, err := UnmarshalSignatureHeader(bytes)
	if err != nil {
		panic(err)
	}
	return sighdr
}
