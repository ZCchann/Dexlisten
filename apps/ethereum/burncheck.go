package ethereum

import (
	"coin/apps/contract"
	"coin/pkg/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)


type information struct {
	name        string   //代币名称
	symbol      string   //代币符号
	totalSupply *big.Int //发行总量
	burn        bool     //燃烧检查结果
}

//燃烧检测
func burnCheck(address common.Address) (*information,error) {
	var information information
	cont, err := contract.NewIERC20(address, Client) //连接
	if err != nil {
		log.Error("连接合约错误",err)
		return nil, err
	}
	name, err := cont.Name(&bind.CallOpts{})
	if err != nil{
		log.Error("获取代币名称错误",err)
		return nil, err
	}
	symbol, err := cont.Symbol(&bind.CallOpts{})
	if err != nil{
		log.Error("获取代币symbol错误",err)
		return nil, err
	}
	totalSupply, err := cont.TotalSupply(&bind.CallOpts{})
	if err != nil{
		log.Error("获取代币发行总量错误",err)
		return nil, err
	}

	information.name = name
	information.symbol = symbol
	information.totalSupply = totalSupply

	//下方为燃烧检查
	_, burn := cont.BURNADDRESS(&bind.CallOpts{})
	if burn == nil {
		information.burn = true
	}else {
		information.burn = false
	}
	_, BurnAddress := cont.BurnAddress(&bind.CallOpts{})
	if BurnAddress == nil {
		information.burn = true
	}else {
		information.burn = false
	}
	return &information ,nil
}