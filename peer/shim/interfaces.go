package shim

import pb "fchain/proto"

type Chaincode interface {
	Init(stub ChaincodeStubInterface) pb.Response

	Invoke(stub ChaincodeStubInterface) pb.Response
}

type ChaincodeStubInterface interface {
	GetArgs() [][]byte

	GetStringArgs() []string

	GetFunctionAndParameters() (string, []string)

	GetTxID() string

	GetState(key string) ([]byte, error)

	PutState(key string, value []byte) error

	GetSignedProposal() (*pb.SignedProposal, error)

	GetCreator() ([]byte, error)

	DelState(key string) error
}
