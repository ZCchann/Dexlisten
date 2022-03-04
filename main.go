package main

import (
	"coin/apps/conf"
	"coin/apps/db"
	"coin/apps/ethereum"
	"coin/apps/tasks"
	"coin/pkg/log"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 二次封装 源：web3.eth.get_transaction_receipt
//type get_transaction_receipt struct {
//	blockHash         common.Hash
//	blockNumber       *big.Int
//	contractAddress   common.Address
//	cumulativeGasUsed uint64
//	gasUsed           uint64
//	logs              []*types.Log
//	status            uint64
//	transactionHash   common.Hash
//	transactionIndex  uint
//	Type              uint8
//	/*
//		缺少以下信息无法获取
//		from
//		to
//	*/
//}

//func Get_transaction_receipt(client *ethclient.Client, transfer_hash string) (transaction_data *get_transaction_receipt, err error) {
//	/*
//		二次封装 源：web3.eth.get_transaction_receipt
//		status字段 1表示有效订单 0表示无效订单
//	*/
//
//	var transaction_receipt get_transaction_receipt
//	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(transfer_hash))
//	if err != nil {
//		return nil, err
//	}
//	transaction_receipt.gasUsed = receipt.GasUsed
//	transaction_receipt.Type = receipt.Type
//	transaction_receipt.status = receipt.Status
//	transaction_receipt.logs = receipt.Logs
//	transaction_receipt.blockHash = receipt.BlockHash
//	transaction_receipt.blockNumber = receipt.BlockNumber
//	transaction_receipt.contractAddress = receipt.ContractAddress
//	transaction_receipt.cumulativeGasUsed = receipt.CumulativeGasUsed
//	transaction_receipt.transactionHash = receipt.TxHash
//	transaction_receipt.transactionIndex = receipt.TransactionIndex
//
//	return &transaction_receipt, nil
//}
/*
type information struct {
	name        string   //代币名称
	symbol      string   //代币符号
	totalSupply *big.Int //发行总量
	burn        bool     //燃烧检查结果
}
func callLP(client *ethclient.Client, LPaddress common.Address) (number *big.Int ,err error) {
	ret, err := contract.NewPancakeLP(LPaddress, client)
	if err != nil {
		fmt.Println(err)
	}
	token0, _ := ret.Token0(&bind.CallOpts{})
	token1, _ := ret.Token1(&bind.CallOpts{})
	WBNB := common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
	BUSD := common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56")
	var reserve *big.Int
	switch  {
	case token0 == WBNB:
		res ,err  := ret.GetReserves(&bind.CallOpts{})
		if err != nil{
			return nil, err
		}
		reserve = res.Reserve0
	case token1 == WBNB:
		res ,err  := ret.GetReserves(&bind.CallOpts{})
		if err != nil{
			return nil, err
		}
		reserve = res.Reserve1
	case token0 == BUSD:
		res ,err  := ret.GetReserves(&bind.CallOpts{})
		if err != nil{
			return nil, err
		}
		reserve = res.Reserve0
	case token1 == BUSD:
		res ,err  := ret.GetReserves(&bind.CallOpts{})
		if err != nil{
			return nil, err
		}
		reserve = res.Reserve1
	}
	return reserve,nil
}

//燃烧检测
func burnCheck(client *ethclient.Client, address common.Address) *information {
	var information information
	cont, err := contract.NewIERC20(address, client) //连接
	if err != nil {
		fmt.Println(err)
	}
	name, _ := cont.Name(&bind.CallOpts{})
	symbol, _ := cont.Symbol(&bind.CallOpts{})
	totalSupply, _ := cont.TotalSupply(&bind.CallOpts{})

	information.name = name
	information.symbol = symbol
	information.totalSupply = totalSupply

	//下方为燃烧检查
	_, burn := cont.BURNADDRESS(&bind.CallOpts{})
	if burn == nil {
		information.burn = true
	}
	_, BurnAddress := cont.BurnAddress(&bind.CallOpts{})
	if BurnAddress == nil {
		information.burn = true
	}else {
		information.burn = false
	}
	return &information
}

func Factorylisten(client *ethclient.Client, address string) {
	contractAddress := common.HexToAddress(address) //EIP55加密合约地址
	query := ethereum.FilterQuery{ //创建连接队列
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs) //监听合约交易信息
	if err != nil {
		log.Fatal(err)
	}
	f, _ := os.Open("apps/contract/PancakeSwap_Factory.json") //读取工厂合约ABI文件
	tokenAbi, err := abi.JSON(f)
	//
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog.TxHash)
			var m = make(map[string]interface{})
			err := tokenAbi.UnpackIntoMap(m, "PairCreated", vLog.Data)

			if err != nil {
				log.Println("Failed to unpack")
				continue
			}
			token0 := common.BytesToAddress(vLog.Topics[1].Bytes())
			token1 := common.BytesToAddress(vLog.Topics[2].Bytes())
			lptoken := m["pair"].(common.Address)

			var b *information //保存燃烧检查的结果
			var pair string    //保存交易对的结果
			var ierc20address string
			WBNB := common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
			BUSD := common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56")

			switch {
			case token1 == WBNB && token0 != WBNB:
				b = burnCheck(client, token0)
				pair = "WBNB"
				ierc20address = token0.String()
			case token1 != WBNB && token0 == WBNB:
				b = burnCheck(client, token1)
				pair = "WBNB"
				ierc20address = token1.String()
			case token1 == BUSD && token0 != BUSD:
				b = burnCheck(client, token0)
				pair = "BUSD"
				ierc20address = token0.String()
			case token1 != BUSD && token0 == BUSD:
				b = burnCheck(client, token1)
				pair = "BUSD"
				ierc20address = token1.String()
			}

			name := b.name
			symbol := b.symbol           //代币符号
			totalSupply := b.totalSupply //发行总量
			burn := b.burn               //燃烧检查结果
			var burncheck string
			if burn == true{
				burncheck = "合约带燃烧"
			}else {
				burncheck = "合约不带燃烧"
			}
			abi := abiCheck(ierc20address)
			var checkabi string
			if abi == true {
				checkabi = "合约已开源"
			}else {
				checkabi = "合约未开源"
			}

			var message strings.Builder
			message.WriteString("PancakeSwap添加了新的交易币对:\n")
			message.WriteString("代币名称: " + name + "\n")
			message.WriteString("代币符号: " + symbol + "\n")
			message.WriteString("代币总发行量: " + web3.FromWei(totalSupply, "ether").String() + "\n")
			message.WriteString("交易对: " + pair + "\n")
			message.WriteString("当前流动性池BNB数量: " + "\n")
			message.WriteString("合约源码情况: " + checkabi + burncheck + "\n")
			message.WriteString("合约地址: " + ierc20address + "\n")
			message.WriteString("Pancake LP代币地址: " + lptoken.String() + "\n")
			s3 := message.String()
			fmt.Println(s3)
		}
	}
}

func sendMessage() {
	telebotID := ""
	telebot := "https://api.telegram.org/bot" + telebotID + "/sendMessage"
	channel_id := ""
	message := ""
	resp, err := http.PostForm(telebot,
		url.Values{"chat_id": {channel_id}, "text": {message}})

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

//检查合约是否开源
func abiCheck(address string) (status bool) {
	bscscan := "https://api.bscscan.com/api?module=contract&action=getabi&address="
	apikey := "&apikey="
	keys := "CF6J1E556BUW1R22F9S7G5SSCZMEQQUFB8"
	var buffer bytes.Buffer
	buffer.WriteString(bscscan)
	buffer.WriteString(address)
	buffer.WriteString(apikey)
	buffer.WriteString(keys)
	u := buffer.String()
	//上方为拼接字符串

	resp, _ := http.Get(u)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	_ = json.Unmarshal(body, &result)
	fmt.Println(result["status"])
	if result["status"] == "0" && result["result"] == "Contract source code not verified" {
		return false
	} else if result["status"] == "1" {
		return true
	}

	return
}
*/

func init() {
	log.Info("main.init...")

	file := "./config.json"
	log.Info("加载配置文件:", file)
	conf.Init(file)

	log.Info("连接redis")
	db.ConnectRedis()
	//初始化区块链主网
	ethereum.Init()

}
func main() {
	signalChan := make(chan os.Signal, 1)
	// 捕捉 Ctrl+c 和 kill 信号，写入signalChan
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	log.Info("main.main...")

	go ethereum.Factorylisten()
	log.Info("启动任务模块")
	go tasks.Start()

	// signalChan阻塞进程
	<-signalChan
	close(signalChan)
	stop()
	fmt.Println()
	log.Info("main.main done.")

}

func stop() {
	ethereum.Client.Close()
	db.RDB.Close()
}

