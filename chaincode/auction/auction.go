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
	"time"
)

// SimpleAuction 拍卖会
type SimpleAuction struct {
}

// Config 配置
type Config struct {
	StartTime   string `json:"startTime"`   //开始时间
	AuctionTime int    `json:"auctionTime"` //竞拍时间
}

// Goods 商品信息
type Goods struct {
	Name  string  `json:"name"`  //商品名称
	Price float64 `json:"price"` //价格
	State int     `json:"state"` //状态
	Owner string  `json:"owner"` // 拥有者
}

const (
	WaitState    = 0 //等待
	AuctionState = 1 //拍卖ing
	SettleState  = 2 //结束
)

// Init 初始化
func (s *SimpleAuction) Init(stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)

}

// Invoke 执行
func (s *SimpleAuction) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error
	if fn == "open" {
		result, err = open(stub, args)
	} else if fn == "set" {
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

// open 开启拍卖
func open(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Excepting start time and auction time")
	}

	startTime := args[0]
	auctionTime, err := strconv.Atoi(args[1])
	if err != nil {
		return "", fmt.Errorf("The auction time should be int")
	}
	config := Config{startTime, auctionTime}
	configStr, err := json.Marshal(config)
	if err != nil {
		return "", fmt.Errorf("Failed to convert config to json")
	}

	err = stub.PutState("config", configStr)
	if err != nil {
		return "", fmt.Errorf("Failed to set config:%s", err.Error())
	}
	return "Success to open auction", nil
}

// set 设置拍卖商品
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 3 {
		return "", fmt.Errorf("Incorrect arguments. Expecting goods name,price,owner")
	}
	name := args[0]
	goodsInfo, err := stub.GetState(name)
	if err != nil {
		return "", fmt.Errorf("Failed to get goods :%s", err.Error())
	} else if goodsInfo != nil {
		return "", fmt.Errorf("This goods already exists :%s", name)
	}

	price, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return "", fmt.Errorf("Failed to convert price to float:%s", err.Error())
	}

	state := WaitState
	owner := args[2]
	goods := Goods{name, price, state, owner}

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

// getConfig 获取配置文件
func getConfig(stub shim.ChaincodeStubInterface) (Config, error) {
	configStr, err := stub.GetState("config")
	config := Config{}

	if err != nil {
		return config, fmt.Errorf("Failed to get auction config")
	}

	json.Unmarshal(configStr, &config)

	return config, nil
}

// checkAuctionStart 检查拍卖是否开启
func checkAuctionStart(startTime string, auctionTime int) bool {
	now := time.Now().Unix()
	timeLayout := "2006-01-02 15:04:05"
	theTime, _ := time.Parse(timeLayout, startTime)

	startTimeUnix := theTime.Unix()

	if now >= startTimeUnix && now < startTimeUnix+int64(auctionTime) {
		return true
	}

	return false
}

// get 获取竞拍商品信息
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a goods name")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("Failed to get goods:%s", err.Error())
	} else if value == nil {
		return "", fmt.Errorf("This goods: %s, not exists", args[0])
	}

	return string(value), nil
}

// start 开始
func start(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	config, err := getConfig(stub)
	if err != nil {
		return "", err
	}

	if checkAuctionStart(config.StartTime, config.AuctionTime) == false {
		return "", fmt.Errorf("The auction isn't on start time")
	}

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
