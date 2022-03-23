package conf

import (
	"coin/pkg/config"
	"coin/pkg/log"
)

type (
	conf struct {
		WBNB     string   `json:"wbnb"         desc:"WBNB"`
		Busd     string   `json:"busd"         desc:"BUSD"`
		Usdc     string   `json:"usdc"         desc:"usdc"`
		Usdt     string   `json:"usdt"         desc:"usdt"`
		ScanApi  string   `json:"scan_api"     desc:"区块链浏览器Api地址"`
		BscRPC   string   `json:"bsc"          desc:"BSC主网链接"`
		Telegram telegram `json:"telegram"     desc:"telegram"`
		Contract contract `json:"contract"     desc:"contract"`
		Redis    redis    `json:"redis"        desc:"redis"`
	}
	contract struct {
		Pancake string `json:"pancake"      desc:"PancakeSwap: Pancake v2合约地址"`
		Biswap  string `json:"biswap"       desc:"Biswap: Factory 合约地址"`
		ApeSwap string `json:"apeswap"       desc:"ApeSwap: ApeFactory 合约地址"`
	}
	redis struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		Database int    `json:"database"`
	}
	telegram struct {
		Botid     string `json:"bot_id"      desc:"telegran bot id"`
		Channelid string `json:"channel_id"  desc:"telegran channel id"`
	}
)

var c = new(conf)

func Init(file string) {
	err := config.BindJSON(file, &c)
	if err != nil {
		log.Error(err)
		panic(err)
	}
}

func Conf() *conf {
	return c
}
