package main

import (
	"errors"
	"time"
)

const (
	slotNum      = 3600
	tickDuration = time.Second
)

type DelayMessage struct {
	curIndex  int
	slots     [slotNum]map[string]*Task
	closed    chan struct{}
	taskClose chan struct{}
	timeClose chan struct{}
	startTime time.Time
}

type TaskFunc func(args ...interface{})

type Task struct {
	cycleNum int
	fun      TaskFunc
	params   []interface{}
}

func NewDelayMessage() *DelayMessage {
	dm := &DelayMessage{
		curIndex:  0,
		closed:    make(chan struct{}),
		taskClose: make(chan struct{}),
		timeClose: make(chan struct{}),
		startTime: time.Now(),
	}
	for i := 0; i < slotNum; i++ {
		dm.slots[i] = make(map[string]*Task)
	}
	return dm
}

func (dm *DelayMessage) Start() {
	go dm.taskLoop()
	go dm.timeLoop()
	select {
	case <-dm.closed:
		{
			dm.taskClose <- struct{}{}
			dm.timeClose <- struct{}{}
			break
		}
	}
}

func (dm *DelayMessage) Close() {
	dm.closed <- struct{}{}
}

func (dm *DelayMessage) taskLoop() {
	for {
		select {
		case <-dm.taskClose:
			return
		default:
			tasks := dm.slots[dm.curIndex]
			for k, v := range tasks {
				if v.cycleNum > 0 {
					v.cycleNum--
					continue
				}
				// do task if cycleNum = 0
				go v.fun(v.params...)
				delete(tasks, k)
			}
		}
	}
}

func (dm *DelayMessage) timeLoop() {
	tick := time.NewTicker(tickDuration)

	for {
		select {
		case <-dm.timeClose:
			return
		case <-tick.C:
			dm.curIndex++
			if dm.curIndex >= slotNum {
				dm.curIndex = 0
			}
		}
	}
}

func (dm *DelayMessage) AddTaskAfter(t time.Duration, key string, fun TaskFunc, params []interface{}) error {
	return dm.AddTaskAt(time.Now().Add(t), key, fun, params)
}

func (dm *DelayMessage) AddTaskAt(t time.Time, key string, fun TaskFunc, params []interface{}) error {
	if time.Now().After(t) {
		return errors.New("invalid time")
	}

	subSecond := t.Unix() - dm.startTime.Unix()
	tasks := dm.slots[subSecond%slotNum]
	if _, ok := tasks[key]; ok {
		return errors.New("task key already exists")
	}

	tasks[key] = &Task{
		cycleNum: int(subSecond / slotNum),
		fun:      fun,
		params:   params,
	}
	return nil
}
