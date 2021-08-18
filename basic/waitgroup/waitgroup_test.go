package waitgroup

import (
	"sync"
	"testing"
	"time"
)

func doWithWaitGroup(t *testing.T, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	time.Sleep(time.Second)
	t.Logf("done~~~")
}

// 某个goroutine等待另外一组goroutine完成
func Test_WaitGroup(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doWithWaitGroup(t, wg)
	}
	go func() {
		wg.Wait()
		t.Logf("我被唤醒了～")
	}()
	go func() {
		wg.Wait()
		t.Logf("我被唤醒了～")
	}()
	wg.Wait() // 等待执行完成
	t.Logf("main groutine 被唤醒了～～～")
}
