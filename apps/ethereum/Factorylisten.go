package ethereum

import (
	"coin/apps/conf"
	"coin/apps/db"
	"coin/apps/web3"
	"coin/pkg/log"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"

	"os"
	"time"
)

func Factorylisten() {
	contractAddress := common.HexToAddress(conf.Conf().Factory) //加载工厂合约地址
	query := ethereum.FilterQuery{ //创建连接队列
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := Client.SubscribeFilterLogs(context.Background(), query, logs) //监听合约交易信息
	if err != nil {
		log.Error(err)
	}
	f, _ := os.Open("apps/contract/PancakeSwap_Factory.json") //读取工厂合约ABI文件
	tokenAbi, err := abi.JSON(f)
	//
	if err != nil {
		log.Error(err)

	}

	for {
		select {
		case err := <-sub.Err():
			log.Error(err)
		case vLog := <-logs:
			var m = make(map[string]interface{})
			err := tokenAbi.UnpackIntoMap(m, "PairCreated", vLog.Data)

			if err != nil {
				log.Error("Failed to unpack")
				continue
			}
			token0 := common.BytesToAddress(vLog.Topics[1].Bytes())
			token1 := common.BytesToAddress(vLog.Topics[2].Bytes())
			lptoken := m["pair"].(common.Address)

			var pair string          //保存交易对的结果
			var ierc20address string //保存非价值代币的结果
			WBNB := common.HexToAddress(conf.Conf().WBNB)
			BUSD := common.HexToAddress(conf.Conf().Busd)
			USDT := common.HexToAddress(conf.Conf().Usdt)
			USDC := common.HexToAddress(conf.Conf().Usdc)

			//判断交易对类型
			switch {
			case token1 == WBNB && token0 != WBNB:
				pair = "WBNB"
				ierc20address = token0.String()
			case token1 != WBNB && token0 == WBNB:
				pair = "WBNB"
				ierc20address = token1.String()
			case token1 == BUSD && token0 != BUSD:
				pair = "BUSD"
				ierc20address = token0.String()
			case token1 != BUSD && token0 == BUSD:
				pair = "BUSD"
				ierc20address = token1.String()
			case token1 == USDT && token0 != USDT:
				pair = "USDT"
				ierc20address = token0.String()
			case token1 != USDT && token0 == USDT:
				pair = "USDT"
				ierc20address = token1.String()
			case token1 == USDC && token0 != USDC:
				pair = "USDC"
				ierc20address = token0.String()
			case token1 != USDC && token0 == USDC:
				pair = "USDC"
				ierc20address = token1.String()
			default:
				pair = ""
				ierc20address = ""
			}

			if pair != "" {
				var b *information                                     //保存燃烧检查的结果
				b, err = burnCheck(common.HexToAddress(ierc20address)) //发动代币合约做燃烧检查
				if err != nil {
					log.Error(err)
				}
				name := b.name                                      //代币名称
				symbol := b.symbol                                  //代币符号
				totalSupply := web3.FromWei(b.totalSupply, "ether") //发行总量 单位ether
				burn := b.burn                                      //燃烧检查结果

				if totalSupply.Cmp(decimal.NewFromInt(1000000)) == 1 || totalSupply.Cmp(decimal.NewFromInt(1000000)) == 0 {
					redis_hset(ierc20address, name, symbol, burn, pair, lptoken.String(), totalSupply.String()) //结果保存到redis内
				}

			}

		}
	}
}

func redis_hset(ierc20address string, name string, symbol string, burn bool, pair string, LP_token_address string, totalSupply string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	redisdb := db.RDB.HMSet(ctx, ierc20address, map[string]interface{}{
		"name":             name,
		"symbol":           symbol,
		"burn":             burn,
		"token_address":    ierc20address,
		"pair":             pair,
		"totalSupply":      totalSupply,
		"LP_token_address": LP_token_address,
	})
	if redisdb.Val() == false {
		log.Error(redisdb)
	}
}
