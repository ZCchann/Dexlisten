package tasks

import (
	"coin/pkg/log"
	"fmt"
	"reflect"
	"runtime"
	"time"
)

var exit bool
var running = make(map[string]bool)

// periodicTask 周期任务注册函数
// 默认每天 00:00:00 执行任务，也可指定时间
func periodicTask(f func(), opts *CronOpts) {
	if opts == nil {
		opts = new(CronOpts)
	}
	funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	Tasks = append(Tasks, funcName)
	go func() {
		for {
			if opts.Duration == 0 {
				// 没有给定间隔时间，就按默认每天+指定时间执行
				opts.Duration = time.Hour * 24
				now := time.Now()
				next := now.Add(opts.Duration)
				if opts.Year == 0 {
					opts.Year = next.Year()
				}
				if opts.Month == 0 {
					opts.Month = next.Month()
				}
				if opts.Day == 0 {
					opts.Day = next.Day()
				}
				next = time.Date(opts.Year, opts.Month, opts.Day, opts.Hour, opts.Min, opts.Sec, 0, next.Location())
				t := time.NewTimer(next.Sub(now))

				<-t.C
			} else {
				// 没有指定时间，每隔一段时间执行一次
				t := time.NewTicker(opts.Duration)

				<-t.C
			}

			if exit {
				break
			}

			running[funcName] = true
			log.Info("Received task:", funcName)
			startTime := time.Now().Unix()
			f()
			endTime := time.Now().Unix()
			log.Info(fmt.Sprintf(
				"Task %s successed in %d second",
				funcName,
				endTime-startTime),
			)
			delete(running, funcName)
		}
	}()
}