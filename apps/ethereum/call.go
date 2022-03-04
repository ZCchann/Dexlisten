package ethereum

import (
	"coin/apps/conf"
	"coin/apps/contract"

	"coin/pkg/log"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func CallLP(LPaddress common.Address) (number *big.Int ,err error) {
	config := conf.Conf()
	ret, err := contract.NewPancakeLP(LPaddress, Client)
	if err != nil {
		fmt.Println(err)
	}
	token0, _ := ret.Token0(&bind.CallOpts{})
	token1, _ := ret.Token1(&bind.CallOpts{})
	WBNB := common.HexToAddress(config.WBNB)
	BUSD := common.HexToAddress(config.Busd)
	USDT := common.HexToAddress(config.Usdt)
	USDC := common.HexToAddress(config.Usdc)

	var reserve *big.Int
	switch  {
	case token0 == WBNB:
		res ,err  := ret.GetReserves(&bind.CallOpts{})
		if err != nil{
			log.Error(err)
			return nil, err
		}
		reserve = res.Reserve0
	case token1 == WBNB:
		res ,err  := ret.GetReserves(&bind.CallOpts{})
		if err != nil{
			log.Error(err)
			return nil, err
		}
		reserve = res.Reserve1
	case token0 == BUSD:
		res ,err  := ret.GetReserves(&bind.CallOpts{})
		if err != nil{
			log.Error(err)
			return nil, err
		}
		reserve = res.Reserve0
	case token1 == BUSD:
		res ,err  := ret.GetReserves(&bind.CallOpts{})
		if err != nil{
			log.Error(err)
			return nil, err
		}
		reserve = res.Reserve1
	case token0 == USDT:
		res ,err  := ret.GetReserves(&bind.CallOpts{})
		if err != nil{
			log.Error(err)
			return nil, err
		}
		reserve = res.Reserve0
	case token1 == USDT:
		res ,err  := ret.GetReserves(&bind.CallOpts{})
		if err != nil{
			log.Error(err)
			return nil, err
		}
		reserve = res.Reserve1
	case token0 == USDC:
		res ,err  := ret.GetReserves(&bind.CallOpts{})
		if err != nil{
			log.Error(err)
			return nil, err
		}
		reserve = res.Reserve0
	case token1 == USDC:
		res ,err  := ret.GetReserves(&bind.CallOpts{})
		if err != nil{
			log.Error(err)
			return nil, err
		}
		reserve = res.Reserve1
	default:
		reserve = nil
	}

	return reserve,nil
}
