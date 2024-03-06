package shim

import (
	pb "fchain/proto"
	"fmt"
)

type KvStore struct{}

func (t *KvStore) Init(stub ChaincodeStubInterface) pb.Response {
	for _, arg := range stub.GetArgs() {
		fmt.Println(string(arg))
	}
	return Success(nil)
}

func (t *KvStore) Invoke(stub ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "write":
		return t.write(stub, args)
	case "delete":
		return t.delete(stub, args)
	case "read":
		return t.read(stub, args)
	default:
		return Error(`Invalid invoke function name. Expecting "write", "delete" or "read"`)
	}
}

func (t *KvStore) write(stub ChaincodeStubInterface, args []string) pb.Response {
	var key string
	var value string

	key = args[0]
	value = args[1]

	err := stub.PutState(key, []byte(value))
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to put" + key + "in state " + "\"}"
		return Error(jsonResp)
	}
	return Success(nil)
}

func (t *KvStore) delete(stub ChaincodeStubInterface, args []string) pb.Response {

	var key string

	key = args[0]

	err := stub.DelState(key)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to delete" + key + "in state " + "\"}"
		return Error(jsonResp)
	}
	return Success(nil)
}

func (t *KvStore) read(stub ChaincodeStubInterface, args []string) pb.Response {

	var key string

	key = args[0]

	_, err := stub.GetState(key)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to delete" + key + "in state " + "\"}"
		return Error(jsonResp)
	}
	return Success(nil)
}
