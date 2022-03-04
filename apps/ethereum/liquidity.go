package ethereum

import (
	"coin/apps/db"
	"coin/apps/telegram_bot"
	"coin/apps/web3"
	"coin/pkg/log"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
	"time"
)

func LiquidityETH() {
	keys, err := redis_keys("*")
	if err != nil {
		panic(err)
	}
	for _, i := range keys {
		var lp_number *big.Int //LP池中价格锚定代币的数量
		ret, err := redis_hgetall(i) //取redis的值
		if err != nil {
			panic(err)
		}
		if n, err := CallLP(common.HexToAddress(ret["LP_token_address"])); err == nil && n != nil {
			lp_number = n
			if lp_number.Cmp(big.NewInt(0)) != 0 {  //判断流动性数量是不是大于0
				var burncheck string
				if ret["burn"] == "1" {
					burncheck = "检出此合约含有燃烧地址，可能带有燃烧"
				} else {
					burncheck = "暂未检出合约带燃烧 请自行鉴别合约源码"
				}
				//开源检测
				abi := abiCheck(ret["token_address"])
				var checkabi string
				if abi == true {
					checkabi = "合约已开源"
				} else {
					checkabi = "合约未开源"
				}
				//整合成string
				var message strings.Builder
				message.WriteString("PancakeSwap添加了新的交易币对:\n")
				message.WriteString("代币名称: " + ret["name"] + "\n")
				message.WriteString("代币符号: " + ret["symbol"] + "\n")
				message.WriteString("代币总发行量: " + ret["totalSupply"] + "\n")
				message.WriteString("交易对: " + ret["pair"] + "\n")
				message.WriteString("当前流动性池" + ret["pair"] + "数量: " + web3.FromWei(lp_number, "ether").String() + "\n")
				message.WriteString("合约源码情况: " + checkabi + ", " + burncheck + "\n")
				message.WriteString("合约地址: " + i + "\n")
				message.WriteString("Pancake LP代币地址: " + ret["LP_token_address"] + "\n")
				s3 := message.String()
				telegram_bot.SendMessage(s3) //使用telegram发送信息到频道
				redis_hdel(i) //跑完了 删除信息
			}
		} else {
			redis_hdel(i)
		}

	}

}
func redis_keys(keys string) (list []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.RDB.Keys(ctx, keys).Result()
}
func redis_hgetall(keys string) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.RDB.HGetAll(ctx, keys).Result() //取redis的值
}

func redis_hdel(keys string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	redis := db.RDB.Del(ctx, keys)
	if redis.Val() == 0 {
		log.Error("redis delet error")
	}
}
