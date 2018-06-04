## FeeCoin
一个基本的虚拟货币交易智能合约

### 功能说明
  - 定价:
  该合约初始化时会制定FeeCoin的交易价格

  - 注册:
  可以通过该功能注册用户

  - 购买:
  用户通过账户余额购买FeeCoin

  - 充值:
  对用户的账户余额进行充值

  - 售出:
  售出FeeCoin获得金额

  - 查询:
  获取用户账户的信息，包括余额，FeeCoin数量

  - 查价:
  获取当前FeeCoin的价格


### 操作说明
请先安装 `Hyperledger Fabric`的 `1.1` 版本

  - 启动环境(Terminal 1)
  ```
  cd env
  docker-compose -f docker-compose-simple.yaml up
  ```

  - 编译chaincode(Terminal 2)
  ```
  docker exec -ti chaincode bash
  cd coin
  go build
  CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=mycc:0 ./coin
  ```

  - 使用chaincode(Terminal 3)
  ```
  docker exec -it cli bash
  peer chaincode install -p chaincodedev/chaincode/coin -n mycc -v 0            # 保证当前目录是chaincodedev的父目录

  peer chaincode instantiate -n mycc -v 0 -c '{"Args":["5.00"]}' -C myc         # 制定FeeCoin价格
  peer chaincode invoke -n mycc -c '{"Args":["set", "a", "200"]}' -C myc        # 注册用户
  peer chaincode invoke -n mycc -c '{"Args":["buy", "a", "10"]}' -C myc         # 买入FeeCoin
  peer chaincode invoke -n mycc -c '{"Args":["query", "a"]}' -C myc             # 查询账户
  peer chaincode invoke -n mycc -c '{"Args":["sale", "a", "10"]}' -C myc        # 卖出FeeCoin
  peer chaincode invoke -n mycc -c '{"Args":["recharge", "a", "1000"]}' -C myc  # 充值
  peer chaincode invoke -n mycc -c '{"Args":["getPrice"]}' -C myc               # 查询FeeCoin当前价格
  ```