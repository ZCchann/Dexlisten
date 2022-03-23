package tasks

import (
	"coin/apps/ethereum"
	"coin/pkg/log"
	"fmt"
	"time"
)

var Tasks = make([]string, 0)

func Start() {
	//periodicTask(demo, &CronOpts{Duration: 10 * time.Second})
	periodicTask(ethereum.LiquidityETH,&CronOpts{Duration: 180 * time.Second}) //每3分钟执行一次
}

func Stop() {
	log.Info(fmt.Sprintf("Stop task [%d]...", len(running)))
	exit = true

	for {
		if len(running) == 0 {
			break
		}
		time.Sleep(time.Second * 1)
	}
}