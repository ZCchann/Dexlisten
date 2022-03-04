package tasks

import (
	"coin/apps/ethereum"
	"time"
)

var Tasks = make([]string, 0)

func Start() {
	//periodicTask(demo, &CronOpts{Duration: 10 * time.Second})
	periodicTask(ethereum.LiquidityETH,&CronOpts{Duration: 180 * time.Second}) //每3分钟执行一次
}

//func demo() {
//	t := time.Now().Format("2006-01-02 15:04:05")
//	log.Info("task.demo.生产者", t)
//	consumer.Messages <- t
//}
