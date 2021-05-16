package basic

import (
	"fmt"
	"testing"
	"time"
)

const (
	timeFormatter = "2006-01-02 15:04:05"
)

func Test_Time01(t *testing.T) {
	now := time.Now()                     // time.Time 类型
	fmt.Printf("current time: %v\n", now) // 2021-05-15 08:40:30.838982 +0800 CST m=+0.000505209

	nowUnix := time.Now().Unix()              // Unix 时间戳，单位：秒
	fmt.Printf("current time: %d\n", nowUnix) // 1621039310

	nowUnixNano := time.Now().UnixNano()          // Unix 时间戳，单位：纳秒
	fmt.Printf("current time: %d\n", nowUnixNano) // 1621039375708197000
}

func Test_Time02(t *testing.T) {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	min := now.Minute()
	sec := now.Second()
	nanoSec := now.Nanosecond()
	fmt.Printf("year: %d, month: %d, day: %d, hour: %d, min: %d, sec: %d, nanoSec: %d\n", year, month, day, hour, min, sec, nanoSec)
}

// 格式化输出时间戳
func Test_Time03(t *testing.T) {
	now := time.Now().Format(timeFormatter)
	fmt.Printf("now: %s\n", now) // 2021-05-15 09:18:30
}

// 时间戳格式转换
func Test_Time04(t *testing.T) {
	tsSec := time.Now().Unix()
	tsStr := time.Unix(tsSec, 0).Format(timeFormatter)
	fmt.Printf("now: %s\n", tsStr)

	tsStr = "2021-05-15 09:22:44"
	ts, err := time.Parse(timeFormatter, tsStr)
	if err == nil {
		tsSec = ts.Unix()
		fmt.Printf("now: %d\n", tsSec)
	}
}
