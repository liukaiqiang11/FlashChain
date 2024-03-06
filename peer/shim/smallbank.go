package shim

import (
	pb "fchain/proto"
	"fmt"
	"strconv"
)

type SmallBank struct{}

var BALANCE int = 100000
var savingTab string = "saving"
var checkingTab string = "checking"

func (s *SmallBank) Init(stub ChaincodeStubInterface) pb.Response {
	for _, arg := range stub.GetArgs() {
		fmt.Println(string(arg))
	}
	return Success(nil)
}

func (s *SmallBank) Invoke(stub ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "setAccount":
		return s.setAccount(stub, args)
	case "almagate":
		return s.almagate(stub, args)
	case "getBalance":
		return s.getBalance(stub, args)
	case "updateBalance":
		return s.updateBalance(stub, args)
	case "updateSaving":
		return s.updateSaving(stub, args)
	case "sendPayment":
		return s.sendPayment(stub, args)
	case "writeCheck":
		return s.writeCheck(stub, args)
	default:
		return Error(`Invalid invoke function name. Expecting "setAccount", "almagate", "getBalance", "updateSaving", "sendPayment" or "writeCheck"`)
	}
}

func (s *SmallBank) setAccount(stub ChaincodeStubInterface, args []string) pb.Response {

	var account string
	var savingamount string
	var checkingamount string
	//var err error

	if len(args) != 3 {
		return Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	account = args[0]
	savingamount = args[1]
	checkingamount = args[2]

	err := stub.PutState(savingTab+"_"+account, []byte(savingamount))
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to put state for " + account + "\"}"
		return Error(jsonResp)
	}

	err = stub.PutState(checkingTab+"_"+account, []byte(checkingamount))
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to put state for " + account + "\"}"
		return Error(jsonResp)
	}
	return Success(nil)
}

func (s *SmallBank) almagate(stub ChaincodeStubInterface, args []string) pb.Response {

	var from string
	var to string

	if len(args) != 2 {
		return Error("Incorrect number of arguments. Expecting name of the person to almagate")
	}

	from = args[0]
	to = args[1]

	var bal1, bal2 int
	var err error
	bal_str1, err := stub.GetState(savingTab + "_" + from)
	if err != nil {
		bal_str1 = []byte(strconv.Itoa(BALANCE))
	}
	bal_str2, err := stub.GetState(checkingTab + "_" + to)
	if err != nil {
		bal_str2 = []byte(strconv.Itoa(BALANCE))
	}

	bal1, err = strconv.Atoi(string(bal_str1))
	if err != nil {
		bal1 = BALANCE
	}
	bal2, err = strconv.Atoi(string(bal_str2))
	if err != nil {
		bal2 = BALANCE
	}
	bal1 += bal2

	err = stub.PutState(checkingTab+"_"+from, []byte("0"))

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to put state for " + from + "\"}"
		return Error(jsonResp)
	}

	err = stub.PutState(savingTab+"_"+to, []byte(strconv.Itoa(bal1)))

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to put state for " + to + "\"}"
		return Error(jsonResp)
	}

	return Success(nil)
}

func (s *SmallBank) getBalance(stub ChaincodeStubInterface, args []string) pb.Response {

	var account string

	if len(args) != 1 {
		return Error("Incorrect number of arguments. Expecting name of the person to getBalance")
	}

	account = args[0]

	var bal1, bal2 int
	var err error
	bal_str1, err := stub.GetState(savingTab + "_" + account)
	if err != nil {
		bal_str1 = []byte(strconv.Itoa(BALANCE))
	}
	bal_str2, err := stub.GetState(checkingTab + "_" + account)
	if err != nil {
		bal_str2 = []byte(strconv.Itoa(BALANCE))
	}

	bal1, err = strconv.Atoi(string(bal_str1))
	if err != nil {
		bal1 = BALANCE
	}
	bal2, err = strconv.Atoi(string(bal_str2))
	if err != nil {
		bal2 = BALANCE
	}
	bal1 += bal2

	//jsonResp := "{\"Key\":\"" + account + "\",\"Value\":\"" + strconv.Itoa(bal1) + "\"}"
	//fmt.Printf("GetBalance Response:%s\n", jsonResp)

	return Success(nil)
}

func (s *SmallBank) updateBalance(stub ChaincodeStubInterface, args []string) pb.Response {
	var account string
	var amountArg string

	if len(args) != 2 {
		return Error("Incorrect number of arguments. Expecting name of the person to updateBalance")
	}

	account = args[0]
	amountArg = args[1]

	bal_str, err2 := stub.GetState(checkingTab + "_" + account)
	if err2 != nil {
		bal_str = []byte(strconv.Itoa(BALANCE))
	}

	var bal1, bal2 int
	var err error
	bal1, err = strconv.Atoi(string(bal_str))
	if err != nil {
		bal1 = BALANCE
	}
	bal2, err = strconv.Atoi(amountArg)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get Balance for " + account + "\"}"
		return Error(jsonResp)
	}
	bal1 += bal2
	err = stub.PutState(checkingTab+"_"+account, []byte(strconv.Itoa(bal1)))

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to put state for " + account + "\"}"
		return Error(jsonResp)
	}

	return Success(nil)
}

func (s *SmallBank) updateSaving(stub ChaincodeStubInterface, args []string) pb.Response {

	var account string
	var amountArg string

	if len(args) != 2 {
		return Error("Incorrect number of arguments. Expecting name of the person to updateSaving")
	}

	account = args[0]
	amountArg = args[1]

	bal_str3, err3 := stub.GetState(savingTab + "_" + account)
	if err3 != nil {
		bal_str3 = []byte(strconv.Itoa(BALANCE))
	}
	var bal1, bal2 int
	var err error

	bal1, err = strconv.Atoi(string(bal_str3))
	if err != nil {
		bal1 = BALANCE
	}
	bal2, err = strconv.Atoi(amountArg)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get Balance for " + account + "\"}"
		return Error(jsonResp)
	}
	bal1 += bal2
	err = stub.PutState(savingTab+"_"+account, []byte(strconv.Itoa(bal1)))

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to put state for " + account + "\"}"
		return Error(jsonResp)
	}

	return Success(nil)
}

func (s *SmallBank) sendPayment(stub ChaincodeStubInterface, args []string) pb.Response {

	var from string
	var to string
	var amountArg string

	if len(args) != 3 {
		return Error("Incorrect number of arguments. Expecting name of the person to sendPayment")
	}

	from = args[0]
	to = args[1]
	amountArg = args[2]

	var bal1, bal2, amount int
	var err error

	bal_str1, err := stub.GetState(checkingTab + "_" + from)
	if err != nil {
		bal_str1 = []byte(strconv.Itoa(BALANCE))
	}
	bal_str2, err := stub.GetState(checkingTab + "_" + to)
	if err != nil {
		bal_str2 = []byte(strconv.Itoa(BALANCE))
	}
	amount, err = strconv.Atoi(amountArg)

	bal1, err = strconv.Atoi(string(bal_str1))
	if err != nil {
		bal1 = BALANCE
	}
	bal2, err = strconv.Atoi(string(bal_str2))
	if err != nil {
		bal2 = BALANCE
	}
	bal1 -= amount
	bal2 += amount

	err = stub.PutState(checkingTab+"_"+from, []byte(strconv.Itoa(bal1)))

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to put state for " + to + "\"}"
		return Error(jsonResp)
	}

	err = stub.PutState(checkingTab+"_"+to, []byte(strconv.Itoa(bal2)))

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to put state for " + to + "\"}"
		return Error(jsonResp)
	}

	return Success(nil)
}

func (s *SmallBank) writeCheck(stub ChaincodeStubInterface, args []string) pb.Response {

	var from string
	var amountArg string

	if len(args) != 2 {
		return Error("Incorrect number of arguments. Expecting name of the person to writeCheck")
	}

	from = args[0]
	amountArg = args[1]

	bal_str2, err2 := stub.GetState(checkingTab + "_" + from)
	if err2 != nil {
		bal_str2 = []byte(strconv.Itoa(BALANCE))
	}
	bal_str3, err3 := stub.GetState(savingTab + "_" + from)
	if err3 != nil {
		bal_str3 = []byte(strconv.Itoa(BALANCE))
	}

	var bal1, bal2 int
	var err error
	var amount int
	bal1, err = strconv.Atoi(string(bal_str2))
	if err != nil {
		bal1 = BALANCE
	}
	bal2, err = strconv.Atoi(string(bal_str3))
	if err != nil {
		bal2 = BALANCE
	}
	amount, err = strconv.Atoi(amountArg)
	if amount < bal1+bal2 {
		err = stub.PutState(checkingTab+"_"+from, []byte(strconv.Itoa(bal1-amount-1)))
	} else {
		err = stub.PutState(checkingTab+"_"+from, []byte(strconv.Itoa(bal1-amount)))
	}

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to put state " + "\"}"
		return Error(jsonResp)
	}

	return Success(nil)
}
