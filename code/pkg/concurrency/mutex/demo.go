package mutex

import (
	"sync"
)

// 线程安全的计数器
type Counter struct {
	mu    sync.Mutex
	value uint64
}

// 加1，线程安全
func (c *Counter) Incr() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// 获取计数器的值，线程安全
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func safeCount() uint64 {
	var c Counter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				c.Incr()
			}
		}()
	}
	wg.Wait()
	return c.Count()
}
