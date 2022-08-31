package message

import (
	"errors"
	"sync"
	"time"
)

const (
	slotNum      = 3600
	tickDuration = time.Second
)

type DelayMessage struct {
	curIndex int
	slots    [slotNum]sync.Map

	startTime time.Time
	ticker    *time.Ticker
	closed    chan struct{}
}

type TaskFunc func(args []any)

type Task struct {
	cycleNum int
	fun      TaskFunc
	params   []any
}

func NewDelayMessage() *DelayMessage {
	dm := &DelayMessage{
		curIndex:  0,
		closed:    make(chan struct{}),
		startTime: time.Now(),
		ticker:    time.NewTicker(tickDuration),
	}
	for i := 0; i < slotNum; i++ {
		dm.slots[i] = sync.Map{}
	}
	return dm
}

func (dm *DelayMessage) Start() {
	loopFunc := func() {
		dm.slots[dm.curIndex].Range(func(k, v any) bool {
			t, ok := v.(*Task)
			if !ok {
				// direct remove if not *Task type
				dm.slots[dm.curIndex].Delete(k)
				return true
			}
			if t.cycleNum > 0 {
				t.cycleNum--
				return true
			}
			// do task in new goroutine
			go t.fun(t.params)
			dm.slots[dm.curIndex].Delete(k)
			return true
		})
	}

	for {
		select {
		case <-dm.closed:
			return
		case <-dm.ticker.C:
			dm.curIndex++
			if dm.curIndex >= slotNum {
				dm.curIndex = 0
			}
			loopFunc()
		}
	}
}

func (dm *DelayMessage) Close() {
	dm.closed <- struct{}{}
	dm.ticker.Stop()
}

func (dm *DelayMessage) AddTaskAfter(t time.Duration, key string, fun TaskFunc, params []any) error {
	return dm.AddTaskAt(time.Now().Add(t), key, fun, params)
}

func (dm *DelayMessage) AddTaskAt(t time.Time, key string, fun TaskFunc, params []any) error {
	if time.Now().After(t) {
		return errors.New("invalid time")
	}

	subSecond := t.Unix() - dm.startTime.Unix()
	cycleNum := int(subSecond / slotNum)

	slotIndex := int(subSecond % slotNum)
	// fast check
	if slotIndex == dm.curIndex && cycleNum == 0 {
		go fun(params)
		return nil
	}

	if _, ok := dm.slots[slotIndex].Load(key); ok {
		return errors.New("task key already exists")
	}

	dm.slots[slotIndex].Store(key, &Task{
		cycleNum: int(subSecond / slotNum),
		fun:      fun,
		params:   params,
	})
	return nil
}
