package ethereum

import (
	"coin/apps/conf"
	"coin/pkg/log"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client

func Init() {
	//连接区块链主网 https://data-seed-prebsc-1-s1.binance.org:8545/
	var err error
	rawUrl := conf.Conf().BscRPC
	Client, err = ethclient.Dial(rawUrl)
	if err != nil {
		log.Error("连接区块链主网失败", err)
		panic(err)
	}
}

