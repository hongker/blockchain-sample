/**
 * 合约币
 */
package main

import (
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

// Init 初始化
func (c *Coin) Init(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	c.Name = args[0]
	c.Creator, _ = stub.GetCreator()

	return shim.Success([]byte("Created success"))

}

// Invoke 调用方法
func (c *Coin) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	if fn == "set" {
		return c.set(stub, args)
	} else if fn == "get" {
		return c.get(stub, args)
	} else if fn == "send" {
		return c.send(stub, args)
	}

	return shim.Error("Function is undefined")
}

// set 设置货币数量
func (c *Coin) set(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting a address and a number")
	}

	_, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("The number must be int")
	}

	user, err := stub.GetCreator()
	if err != nil {
		return shim.Error(fmt.Sprintf("Can't get the user:", err.Error()))
	}

	if string(user) != string(c.Creator) { // 判断是否是创造者
		return shim.Error("Permission denied")
	}

	stub.PutState(args[0], []byte(args[1]))

	return shim.Success([]byte("Set success"))
}

// get 获取货币数量
func (c *Coin) get(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting a address")
	}

	number, _ := stub.GetState(args[0])
	return shim.Success(number)
}

// send 发送货币
func (c *Coin) send(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting a address and a number")
	}

	tradeNumber, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("The number must be int")
	}

	sender, _ := stub.GetCreator() //发送方

	senderNumberByte, _ := stub.GetState(string(sender))
	receiverNumberByte, err := stub.GetState(args[0]) //接收方

	if err != nil {
		return shim.Error("The user is not exist")
	}

	// convert byte to int
	senderNumber, _ := strconv.Atoi(string(senderNumberByte))
	receiverNumber, _ := strconv.Atoi(string(receiverNumberByte))

	if senderNumber < tradeNumber {
		return shim.Error("Your coin is not enough")
	}

	stub.PutState(string(sender), []byte(strconv.Itoa(senderNumber-tradeNumber)))
	stub.PutState(args[0], []byte(strconv.Itoa(receiverNumber+tradeNumber)))

	return shim.Success([]byte("Send success"))

}

// main 主函数
func main() {
	if err := shim.Start(new(Coin)); err != nil {
		fmt.Printf("Error starting Coin chaincode: %s", err)
	}
}
