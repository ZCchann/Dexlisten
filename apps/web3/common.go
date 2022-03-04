package web3

import (
	"github.com/shopspring/decimal"
	"math/big"
)


func toWei(iamount interface{}, decimals int) *big.Int {
	// ToWei decimals to wei
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

func ToWei(number interface{}, Uint string) *big.Int {
	switch {
	case Uint == "wei":
		//n := decimal.NewFromFloat(number * 1)
		//r := new(big.Int)
		//r.SetString(n.String(), 10)
		return toWei(number, 0)

	case Uint == "kwei", Uint == "babbage", Uint == "femtoether":
		//n := decimal.NewFromFloat(number * 1000)
		//r := new(big.Int)
		//r.SetString(n.String(), 10)
		//return r
		return toWei(number, 3)
	case Uint == "mwei", Uint == "lovelace", Uint == "picoether":
		//n := decimal.NewFromFloat(number * 1000000)
		//r := new(big.Int)
		//r.SetString(n.String(), 10)
		//return r
		return toWei(number, 6)
	case Uint == "gwei", Uint == "shannon", Uint == "nanoether", Uint == "nano":
		//n := decimal.NewFromFloat(number * 1000000000)
		//r := new(big.Int)
		//r.SetString(n.String(), 10)
		//return r
		return toWei(number, 9)
	case Uint == "szabo", Uint == "microether", Uint == "micro":
		//n := decimal.NewFromFloat(number * 1000000000000)
		//r := new(big.Int)
		//r.SetString(n.String(), 10)
		//return r
		return toWei(number, 12)
	case Uint == "finney", Uint == "milliether", Uint == "milli":
		//n := decimal.NewFromFloat(number * 1000000000000000)
		//r := new(big.Int)
		//r.SetString(n.String(), 10)
		//return r
		return toWei(number, 15)
	case Uint == "ether":
		//n := decimal.NewFromFloat(number * 1000000000000000000)
		//r := new(big.Int)
		//r.SetString(n.String(), 10)
		//return r
		return toWei(number, 18)
	case Uint == "grand", Uint == "kether":
		//n := decimal.NewFromFloat(number * 1000000000000000000000)
		//r := new(big.Int)
		//r.SetString(n.String(), 10)
		//return r
		return toWei(number, 21)
	case Uint == "mether":
		//n := decimal.NewFromFloat(number * 1000000000000000000000000)
		//r := new(big.Int)
		//r.SetString(n.String(), 10)
		//return r
		return toWei(number, 24)
	case Uint == "gether":
		//n := decimal.NewFromFloat(number * 1000000000000000000000000000)
		//r := new(big.Int)
		//r.SetString(n.String(), 10)
		//return r
		return toWei(number, 27)
	case Uint == "tether":
		//n := decimal.NewFromFloat(number * 1000000000000000000000000000000)
		//r := new(big.Int)
		//r.SetString(n.String(), 10)
		//return r
		return toWei(number, 30)
	default:
		return big.NewInt(0)
	}

}

func toDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	// ToDecimal wei to decimals
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

func FromWei(number interface{}, Uint string) decimal.Decimal {

	switch {
	case Uint == "wei":
		n := toDecimal(number, 0)
		return n
	case Uint == "kwei", Uint == "babbage", Uint == "femtoether":
		return toDecimal(number, 3)
	case Uint == "mwei", Uint == "lovelace", Uint == "picoether":
		return toDecimal(number, 6)
	case Uint == "gwei", Uint == "shannon", Uint == "nanoether", Uint == "nano":
		return toDecimal(number, 9)
	case Uint == "szabo", Uint == "microether", Uint == "micro":
		return toDecimal(number, 12)
	case Uint == "finney", Uint == "milliether", Uint == "milli":
		return toDecimal(number, 15)
	case Uint == "ether":
		return toDecimal(number, 18)
	case Uint == "grand", Uint == "kether":
		return toDecimal(number, 21)
	case Uint == "mether":
		return toDecimal(number, 24)
	case Uint == "gether":
		return toDecimal(number, 27)
	case Uint == "tether":
		return toDecimal(number, 30)
	//case Uint == "kwei", Uint == "babbage", Uint == "femtoether":
	//	n := decimal.NewFromFloat(number / float64(1000))
	//	return n
	//
	//case Uint == "mwei", Uint == "lovelace", Uint == "picoether":
	//	n := decimal.NewFromFloat(number / float64(1000000))
	//	return n
	//case Uint == "gwei", Uint == "shannon", Uint == "nanoether", Uint == "nano":
	//	n := decimal.NewFromFloat(number / float64(1000000000))
	//
	//	return n
	//case Uint == "szabo", Uint == "microether", Uint == "micro":
	//	n := decimal.NewFromFloat(number / float64(1000000000000))
	//
	//	return n
	//case Uint == "finney", Uint == "milliether", Uint == "milli":
	//	n := decimal.NewFromFloat(number / float64(1000000000000000))
	//
	//	return n
	//case Uint == "ether":
	//	n := decimal.NewFromFloat(number / float64(1000000000000000000))
	//
	//	return n
	//case Uint == "grand", Uint == "kether":
	//	n := decimal.NewFromFloat(number / float64(1000000000000000000000))
	//
	//	return n
	//case Uint == "mether":
	//	n := decimal.NewFromFloat(number / float64(1000000000000000000000000))
	//
	//	return n
	//case Uint == "gether":
	//	n := decimal.NewFromFloat(number / float64(1000000000000000000000000000))
	//	return n
	//
	//case Uint == "tether":
	//	n := decimal.NewFromFloat(number / float64(1000000000000000000000000000000))
	//	return n

	default:
		return decimal.NewFromFloat(0)
	}
}

