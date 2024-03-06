// Copyright the Hyperledger Fabric contributors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package shim

import (
	"fchain/proto/util"
	"fmt"

	pb "fchain/proto"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

type ChaincodeStub struct {
	TxID           string
	handler        *Handler
	args           [][]byte
	signedProposal *pb.SignedProposal
	proposal       *pb.Proposal
	creator        []byte
}

func newChaincodeStub(handler *Handler, txid string, input *pb.ChaincodeInput, signedProposal *pb.SignedProposal) (*ChaincodeStub, error) {
	stub := &ChaincodeStub{
		handler:        handler,
		TxID:           txid,
		args:           input.Args,
		signedProposal: signedProposal,
	}
	if signedProposal != nil {
		var err error

		stub.proposal = &pb.Proposal{}
		err = proto.Unmarshal(signedProposal.ProposalBytes, stub.proposal)
		if err != nil {
			return nil, fmt.Errorf("failed to extract Proposal from SignedProposal: %s", err)
		}

		if len(stub.proposal.GetHeader()) == 0 {
			return nil, errors.New("failed to extract Proposal fields: proposal header is nil")
		}

		shdr := &pb.SignatureHeader{}
		if err := proto.Unmarshal(stub.proposal.GetHeader(), shdr); err != nil {
			return nil, fmt.Errorf("failed to extract signature header: %s", err)
		}
		stub.creator = shdr.GetCreator()

		payload := &pb.ChaincodeSpec{}
		if err := proto.Unmarshal(stub.proposal.GetPayload(), payload); err != nil {
			return nil, fmt.Errorf("failed to extract proposal payload: %s", err)
		}

	}

	return stub, nil
}

func (s *ChaincodeStub) GetTxID() string {
	return s.TxID
}

func (s *ChaincodeStub) GetArgs() [][]byte {
	return s.args
}

func (s *ChaincodeStub) GetStringArgs() []string {
	args := s.GetArgs()
	strargs := make([]string, 0, len(args))
	for _, barg := range args {
		strargs = append(strargs, string(barg))
	}
	return strargs
}

func (s *ChaincodeStub) GetFunctionAndParameters() (function string, params []string) {
	allargs := s.GetStringArgs()
	function = ""
	params = []string{}
	if len(allargs) >= 1 {
		function = allargs[0]
		params = allargs[1:]
	}
	return
}

func (s *ChaincodeStub) GetState(key string) ([]byte, error) {
	if key == "" {
		return nil, errors.New("key must not be an empty string")
	}
	return s.handler.handleGetState(key, s.TxID)
}

func (s *ChaincodeStub) PutState(key string, value []byte) error {
	if key == "" {
		return errors.New("key must not be an empty string")
	}

	return s.handler.handlePutState(key, value, s.TxID)
}

func (s *ChaincodeStub) DelState(key string) error {
	if key == "" {
		return errors.New("key must not be an empty string")
	}

	return s.handler.handleDelState(key, s.TxID)
}

func (s *ChaincodeStub) GetSignedProposal() (*pb.SignedProposal, error) {
	return s.signedProposal, nil
}

// GetProposal 获取交易提案
func (s *ChaincodeStub) GetProposal() (*pb.Proposal, error) {
	return s.proposal, nil
}

// GetCreator 获取交易创建者
func (s *ChaincodeStub) GetCreator() ([]byte, error) {
	return s.creator, nil
}

func (s *ChaincodeStub) getChainCodeName() (string, error) {
	if s.proposal == nil {
		return "", errors.New("ChaincodeStub proposal is nil")
	}
	if s.proposal.Payload == nil {
		return "", errors.New("ChaincodeStub proposal Payload is nil")
	}
	cisByte := s.proposal.Payload
	spec, err := util.UnmarshalChaincodeSpec(cisByte)
	if err != nil {
		return "", err
	}
	if spec.ChaincodeID == nil {
		return "", errors.New("spec ChaincodeID is nil")
	}

	return spec.ChaincodeID.Name, nil
}
