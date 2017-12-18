/**
 * 智能合约-拍卖
 *
 * @author hongker
 * @date 2017-12-16
 */
package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

// SimpleAuction 拍卖会
type SimpleAuction struct {
	StartTime   int64 //开始时间
	AuctionTime int   //竞拍时间
}

type Goods struct {
	Price float64 `json:"price"` //价格
	State int     `json:"state"` //状态
}

const WaitState = 0    //等待
const AuctionState = 1 //拍卖ing
const SettleState = 2  //结束

// Init 初始化
func (s *SimpleAuction) Init(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("Incorrect params!Excepting StartTime and AuctionTime")
	}
	var err error

	s.StartTime, err = strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return shim.Error("StartTime should be int")
	}

	s.AuctionTime, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("AuctionTime should be int ")
	}

	return shim.Success(nil)

}

// Invoke 执行
func (s *SimpleAuction) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error
	if fn == "set" {
		result, err = set(stub, args)
	} else if fn == "get" {
		result, err = get(stub, args)
	} else if fn == "start" {
		result, err = start(stub, args)
	} else if fn == "auction" {
		result, err = auction(stub, args)
	} else if fn == "settle" {
		result, err = settle(stub, args)
	}

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(result))
}

// set 设置拍卖商品
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a goods name and a price")
	}

	price, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return "", fmt.Errorf("Failed to convert price to float:%s", err.Error())
	}

	state := WaitState
	goods := Goods{price, state}

	goodsStr, err := json.Marshal(goods)
	if err != nil {
		return "", fmt.Errorf("Failed to convert goods to string:%s", err.Error())
	}

	err = stub.PutState(args[0], goodsStr)
	if err != nil {
		return "", fmt.Errorf("Failed to set auction goods:%s", err.Error())
	}

	return fmt.Sprintf("Set auction Success"), nil
}

// get 获取竞拍商品信息
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a goods name")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get goods price:%s, with error:%s", args[0], err.Error())
	}

	return string(value), nil
}

// start 开始
func start(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get goods info:%s, with error:%s", args[0], err.Error())
	}

	goods := Goods{}
	json.Unmarshal(value, &goods)
	if goods.State != WaitState {
		return "", fmt.Errorf("The goods state is :%s, can't start", goods.State)
	}

	goods.State = AuctionState
	goodsStr, err := json.Marshal(goods)
	if err != nil {
		return "", fmt.Errorf("Failed to convert goods to string:%s", err.Error())
	}

	err = stub.PutState(args[0], goodsStr)
	if err != nil {
		return "", fmt.Errorf("Failed to start:%s", err.Error())
	}

	return fmt.Sprintf("Start Success"), nil

}

// auction 出价
func auction(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a goods name and a price")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get goods price:%s, with error:%s", args[0], err.Error())
	}

	goods := Goods{}
	json.Unmarshal(value, &goods)
	if goods.State != AuctionState {
		return "", fmt.Errorf("The goods state is :%s, can't auction", goods.State)
	}

	auctionPrice, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return "", fmt.Errorf("Failed to convert price to float:%s", err.Error())
	}

	if auctionPrice <= goods.Price {
		return "", fmt.Errorf("Your price: %s should be higher than current price: %s", args[1], goods.Price)
	}

	goods.Price = auctionPrice
	goodsStr, err := json.Marshal(goods)
	if err != nil {
		return "", fmt.Errorf("Failed to convert goods to string:%s", err.Error())
	}

	err = stub.PutState(args[0], goodsStr)
	if err != nil {
		return "", fmt.Errorf("Failed to auction:%s", err.Error())
	}

	return fmt.Sprintf("Auction Success"), nil
}

// settle 结算
func settle(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a goods name")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get goods info:%s, with error:%s", args[0], err.Error())
	}

	goods := Goods{}
	json.Unmarshal(value, &goods)
	if goods.State != AuctionState {
		return "", fmt.Errorf("The goods state is :%s, can't settle", goods.State)
	}

	goods.State = SettleState
	goodsStr, err := json.Marshal(goods)
	if err != nil {
		return "", fmt.Errorf("Failed to convert goods to string:%s", err.Error())
	}

	err = stub.PutState(args[0], goodsStr)
	if err != nil {
		return "", fmt.Errorf("Failed to settle:%s", err.Error())
	}

	return fmt.Sprintf("Settle Success"), nil

}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAuction)); err != nil {
		fmt.Printf("Error starting SimpleAuction chaincode: %s", err)
	}
}
