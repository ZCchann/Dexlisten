package tasks

import "time"

type CronOpts struct {
	Year     int           `remark:"指定到年份"`
	Month    time.Month    `remark:"指定到月份"`
	Day      int           `remark:"指定到日期"`
	Hour     int           `remark:"指定到小时"`
	Min      int           `remark:"指定到分钟"`
	Sec      int           `remark:"指定到秒"`
	Duration time.Duration `remark:"间隔时间，如果给定了该时间，以上时间参数将无效；如不给定，则为每天的以上时间参数为主；"`
}
