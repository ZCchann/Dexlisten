package telegram_bot

import (
	"coin/apps/conf"
	"coin/pkg/log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendMessage(message string) {
	telebot := "https://api.telegram.org/bot" + conf.Conf().Telegram.Botid + "/sendMessage" //拼接url
	resp, err := http.PostForm(telebot,
		url.Values{"chat_id": {conf.Conf().Telegram.Channelid}, "text": {message}})

	if err != nil {
		log.Error(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return
	}
	m := make(map[string]interface{})  //创建一个map接收post的返回值
	_ = json.Unmarshal(body, &m)
	if m["ok"] == "false"{ //如果出现不是ok的返回个报错 提供检查
		log.Error("telegram error")
		log.Error(m)
	}
}
