package ethereum

import (
	"bytes"
	"coin/apps/conf"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//检查合约是否开源
func abiCheck(address string) (status bool) {
	bscscan := "https://api.bscscan.com/api?module=contract&action=getabi&address="
	apikey := "&apikey="
	keys := conf.Conf().ScanApi
	var buffer bytes.Buffer
	buffer.WriteString(bscscan)
	buffer.WriteString(address)
	buffer.WriteString(apikey)
	buffer.WriteString(keys)
	u := buffer.String()
	//上方为拼接字符串

	resp, _ := http.Get(u) //发送get请求
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	_ = json.Unmarshal(body, &result)
	var s bool
	if result["status"] == "0" && result["result"] == "Contract source code not verified" {
		s = false
	} else if result["status"] == "1" {
		s = true
	}

	return s
}