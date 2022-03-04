# Tasks

## 说明

任务app，主要用于定时执行任务

## 使用方法

在 Start 中注册函数，并给定时间选项

```go
// tasks.go

func Start() {
    periodicTask(demo, &CronOpts{Duration: 10 * time.Second})
}
```

## 任务-1

定时获取打包池的数据
