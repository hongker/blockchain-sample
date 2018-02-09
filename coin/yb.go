package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

// Coin 货币
type Coin struct {
	Name    string //名称
	Creator []byte //创造者
}

// Account 账户
type Account struct {
	ID     int64  // 编号
	Name   string // 用户名
	Amount int64  // 存款,单位:分
	Number int64  // 代币数量
}

// Init 初始化
func (c *Coin) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// Invoke 调用方法
func (c *Coin) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	if fn == "get" {
		return c.get(stub, args)
	} else if fn == "trade" {
		return c.trade(stub, args)
	} else if fn == "send" {
		return c.send(stub, args)
	}

	return shim.Error("The function is undefined")
}

// get 获取货币数量
func (c *Coin) get(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting a username")
	}

	accountVal, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("The account is not exist")
	}
	return shim.Success(accountVal)
}

// trade 交易
func (c *Coin) trade(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting buyer,seller,price,number")
	}
	buyer := args[0]
	seller := args[1]
	price, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		return shim.Error("Incorrect Price")
	}

	number, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		return shim.Error("Incorrect Number")
	}

	amount := price * number

	buyerAccountVal, err := stub.GetState(buyer)
	if err != nil {
		return shim.Error("The buyer not exist")
	}
	buyerAccount := Account{}
	json.Unmarshal(buyerAccountVal, &buyerAccount)

	sellerAccountVal, err := stub.GetState(seller)
	if err != nil {
		return shim.Error("The buyer not exist")
	}
	sellerAccount := Account{}
	json.Unmarshal(sellerAccountVal, &sellerAccount)

	if buyerAccount.Amount < amount {
		return shim.Error("Your amount is not enough")
	}
	// 扣钱,得币
	buyerAccount.Number = buyerAccount.Number + number
	buyerAccount.Amount = buyerAccount.Amount - amount

	if sellerAccount.Number < number {
		return shim.Error("Your coin number is not enough")
	}
	// 扣币,得钱
	sellerAccount.Number = sellerAccount.Number - number
	sellerAccount.Amount = sellerAccount.Amount + amount

	buyerAccountVal, err = json.Marshal(buyerAccount)
	if err != nil {
		return shim.Error("Failed to convert buyer account to string" + err.Error())
	}

	sellerAccountVal, err = json.Marshal(sellerAccount)
	if err != nil {
		return shim.Error("Failed to convert seller account to string" + err.Error())
	}

	err = stub.PutState(buyer, buyerAccountVal)
	if err != nil {
		return shim.Error("Failed to update buyer info:" + err.Error())
	}

	err = stub.PutState(seller, sellerAccountVal)
	if err != nil {
		return shim.Error("Failed to update seller info:" + err.Error())
	}
	return shim.Success(nil)
}

// send 转账(type:1-转钱,2-转币)
func (c *Coin) send(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting sender,accepter,amount,type")
	}
	sender := args[0]
	acceptor := args[1]
	amount, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		return shim.Error("Incorrect amount")
	}

	senderAccountVal, err := stub.GetState(sender)
	if err != nil {
		return shim.Error("The sender is not exist")
	}
	senderAccount := Account{}
	json.Unmarshal(senderAccountVal, &senderAccount)

	acceptorAccountVal, err := stub.GetState(acceptor)
	if err != nil {
		return shim.Error("The acceptor not exist")
	}
	acceptorAccount := Account{}
	json.Unmarshal(acceptorAccountVal, &acceptorAccount)

	if args[3] == "1" { // 转钱
		if senderAccount.Amount < amount {
			return shim.Error("Your amount is not enough")
		}
		// 扣钱
		senderAccount.Amount = senderAccount.Amount - amount

		// 得钱
		acceptorAccount.Amount = acceptorAccount.Amount + amount
	} else { //转币
		if senderAccount.Number < amount {
			return shim.Error("Your number is not enough")
		}
		// 扣币
		senderAccount.Number = senderAccount.Number - amount

		// 得币
		acceptorAccount.Number = acceptorAccount.Number + amount
	}

	senderAccountVal, err = json.Marshal(senderAccount)
	if err != nil {
		return shim.Error("Failed to convert sender account to string" + err.Error())
	}

	acceptorAccountVal, err = json.Marshal(acceptorAccount)
	if err != nil {
		return shim.Error("Failed to convert acceptor account to string" + err.Error())
	}

	err = stub.PutState(sender, senderAccountVal)
	if err != nil {
		return shim.Error("Failed to update sender info:" + err.Error())
	}

	err = stub.PutState(acceptor, acceptorAccountVal)
	if err != nil {
		return shim.Error("Failed to update acceptor info:" + err.Error())
	}
	return shim.Success(nil)
}

// main 主函数
func main() {
	if err := shim.Start(new(Coin)); err != nil {
		fmt.Printf("Error starting Coin chaincode: %s", err)
	}
}
