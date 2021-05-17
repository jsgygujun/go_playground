package basic

import (
	"fmt"
	"testing"
	"time"
)

// 设定超时时间
func Test_Timer01(t *testing.T) {
	timer := time.NewTimer(1 * time.Second)
	select {
	case <-timer.C:
		fmt.Printf("timeout!\n")
		//case <-conn:
		//	timer.Stop()
	}
}

// 延迟执行某个方法
func Test_Timer02(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	select {
	case <-timer.C:
		fmt.Printf("delay 5s")
	}
}

// 等待指定的时间, time.After
func Test_Timer03(t *testing.T) {
	fmt.Printf("now: %v\n", time.Now())
	<-time.After(5 * time.Second) // 阻塞等待5秒
	fmt.Printf("new: %v\n", time.Now())
}

// 非阻塞延迟执行, time.AfterFunc
func Test_Timer04(t *testing.T) {
	fmt.Printf("now: %v\n", time.Now())
	time.AfterFunc(5*time.Second, func() {
		fmt.Printf("exec: %v\n", time.Now())
	})
	fmt.Printf("main: %v\n", time.Now())
	time.Sleep(6 * time.Second)
}
