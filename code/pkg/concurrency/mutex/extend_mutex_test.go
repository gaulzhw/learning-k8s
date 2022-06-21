package mutex

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTryLock(t *testing.T) {
	var mu Mutex
	go func() { // 启动一个goroutine持有一段时间的锁
		mu.Lock()
		waitFor := rand.Intn(3)
		fmt.Printf("wait for %d seconds\n", waitFor)
		time.Sleep(time.Duration(waitFor) * time.Second)
		mu.Unlock()
	}()
	time.Sleep(time.Second)
	ok := mu.TryLock() // 尝试获取到锁
	if ok {
		// 获取成功
		fmt.Println("got the lock") // do something
		mu.Unlock()
		return
	}
	// 没有获取到
	fmt.Println("can't get the lock")
}

func TestCount(t *testing.T) {
	var mu Mutex
	for i := 0; i < 1000; i++ { // 启动1000个goroutine
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second) // 输出锁的信息
	fmt.Printf("waitings: %d, isLocked: %t, woken: %t, starving: %t\n", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}
