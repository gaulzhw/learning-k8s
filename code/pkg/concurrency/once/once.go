package once

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

// 传入的函数f有返回值error，如果初始化失败，需要返回失败的error
// Do方法会把这个error返回给调用者
func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		//fast path
		return nil
	}
	return o.slowDo(f)
}

// 如果还没有初始化
func (o *Once) slowDo(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 {
		// 双检查，还没有初始化
		err = f()
		if err == nil {
			// 初始化成功才将标记置为已初始化
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}

// Done 返回此Once是否执行过
// 如果执行过则返回true
// 如果没有执行过或者正在执行，返回false
func (o *Once) Done() bool {
	return atomic.LoadUint32(&o.done) == 1
}
