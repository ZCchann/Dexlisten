package main

import (
	"coin/apps/conf"
	"coin/apps/db"
	"coin/apps/ethereum"
	"coin/apps/tasks"
	"coin/pkg/log"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)



func init() {
	log.Info("main.init...")

	file := "./config.json"
	log.Info("加载配置文件:", file)
	conf.Init(file)

	log.Info("连接redis")
	db.ConnectRedis()
	//初始化区块链主网
	ethereum.Init()

}
func main() {
	signalChan := make(chan os.Signal, 1)
	// 捕捉 Ctrl+c 和 kill 信号，写入signalChan
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	log.Info("main.main...")

	log.Info("启动监控模块")
	ethereum.FactoryInit()
	log.Info("启动任务模块")
	go tasks.Start()

	// signalChan阻塞进程
	<-signalChan
	close(signalChan)
	stop()
	fmt.Println()
	log.Info("main.main done.")

}

func stop() {

	ethereum.Client.Close()
	db.RDB.Close()
}

