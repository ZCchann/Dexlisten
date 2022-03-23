package ethereum

import (
	"coin/apps/conf"
	"coin/pkg/log"
)

func FactoryInit() {
	log.Info("启动PancakeSwap监控")
	go Factorylisten(conf.Conf().Contract.Pancake)
	log.Info("启动Biswap监控")
	go Factorylisten(conf.Conf().Contract.Biswap)
	log.Info("启动ApeSwap监控")
	go Factorylisten(conf.Conf().Contract.ApeSwap)
}