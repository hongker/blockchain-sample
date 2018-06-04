package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"encoding/json"
)

// SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
}

// Account 账户
type Account struct {
	Name string		// 用户名
	Amount float64	// 金额
	Count int64		// 币数
}

func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// Get the args from the transaction proposal
	args := stub.GetStringArgs()
	if len(args) != 1 {
		return shim.Error("Incorrect arguments. Expecting a price")
	}

	price, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return shim.Error("Incorrect arguments of price")
	}

	// 初始化价格
	err = setPrice(stub, price)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error
	if fn == "set" {
		err = set(stub, args)
	} else if fn == "recharge" {
		err = recharge(stub, args)
	} else if fn == "buy" {
		err = buy(stub, args)
	} else if fn == "sale" {
		err = sale(stub, args)
	} else if fn == "query" {
		result, err = query(stub, args)
	} else if fn == "getPrice" {
		price, err := getPrice(stub)
		if err == nil {
			result = fmt.Sprintf("%f", price)
		}

	} else { // assume 'get' even if fn is nil
		err = fmt.Errorf("Func isn't exists")
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	// Return the result as success payload
	return shim.Success([]byte(result))
}

// setPrice 设置价格
func setPrice(stub shim.ChaincodeStubInterface, price float64) error {
	priceStr := fmt.Sprintf("%f", price)
	err := stub.PutState("price", []byte(priceStr))

	if err != nil {
		return fmt.Errorf("Failed to set price with error: %s", err)
	}

	return nil
}

// getPrice 获取当前价格
func getPrice(stub shim.ChaincodeStubInterface) (float64, error) {
	value, err := stub.GetState("price")

	if err != nil {
		return 0, fmt.Errorf("Failed to get price with error: %s", err)
	}

	return strconv.ParseFloat(string(value), 64)
}

// set 初始化
func set(stub shim.ChaincodeStubInterface, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Incorrect arguments. Expecting a name, a amount")
	}

	amountFloat, err := strconv.ParseFloat(args[1], 10)
	if err != nil {
		return fmt.Errorf("Failed to convert float")
	}
	a := Account{args[0], amountFloat, 0}

	accountJson, err := json.Marshal(a)
	if err != nil {
		return fmt.Errorf("Failed to convert json")
	}

	err = stub.PutState(a.Name, []byte(accountJson))

	return err
}

// recharge 充值
func recharge(stub shim.ChaincodeStubInterface, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Incorrect arguments. Expecting a name, a amount")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return fmt.Errorf("Failed to get account: %s with error: %s", args[0], err)
	}

	var a Account
	json.Unmarshal(value, &a)

	amount, err := strconv.ParseFloat(args[1], 10)
	if err != nil {
		return fmt.Errorf("Failed to convert float")
	}

	a.Amount = a.Amount + amount

	accountJson, err := json.Marshal(a)
	if err != nil {
		return fmt.Errorf("Failed to convert json")
	}

	err = stub.PutState(a.Name, []byte(accountJson))

	return err
}

// query 查询
func query(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a name, a count")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}

	return string(value), nil;

}

// buy 购买
func buy(stub shim.ChaincodeStubInterface, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Incorrect arguments. Expecting a name, a count")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return fmt.Errorf("Failed to get account: %s with error: %s", args[0], err)
	}

	var a Account
	json.Unmarshal(value, &a)

	count, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return fmt.Errorf("Failed to convert int")
	}

	if count <= 0 {
		return fmt.Errorf("The count should gt 0")
	}

	price, err := getPrice(stub)
	if err != nil {
		return fmt.Errorf("Failed to get price:" + err.Error())
	}

	amount := float64(count) * price
	if a.Amount < amount {
		return fmt.Errorf("Your dont't have enough money")
	}

	a.Amount = a.Amount - amount
	a.Count = a.Count + count

	accountJson, err := json.Marshal(a)
	if err != nil {
		return fmt.Errorf("Failed to convert json")
	}

	err = stub.PutState(a.Name, []byte(accountJson))

	return err
}

// sale 卖出
func sale(stub shim.ChaincodeStubInterface, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Incorrect arguments. Expecting a name, a count")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}

	var a Account
	json.Unmarshal(value, &a)

	count, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return fmt.Errorf("Failed to convert int")
	}

	if count <= 0 {
		return fmt.Errorf("The count should gt 0")
	}

	if a.Count < count {
		return fmt.Errorf("Your dont't have enough coin")
	}

	price, err := getPrice(stub)
	if err != nil {
		return fmt.Errorf("Failed to get price:" + err.Error())
	}

	amount := float64(count) * price

	a.Amount = a.Amount + amount
	a.Count = a.Count - count

	accountJson, err := json.Marshal(a)
	if err != nil {
		return fmt.Errorf("Failed to convert json")
	}

	err = stub.PutState(a.Name, []byte(accountJson))

	return err
}


func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}