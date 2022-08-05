package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDelayMessage(t *testing.T) {
	dm := NewDelayMessage()

	dm.AddTaskAfter(time.Second*10, "test1", func(args ...interface{}) {
		fmt.Println(args...)
	}, []interface{}{1, 2, 3})
	dm.AddTaskAfter(time.Second*10, "test2", func(args ...interface{}) {
		fmt.Println(args...)
	}, []interface{}{4, 5, 6})
	dm.AddTaskAfter(time.Second*20, "test3", func(args ...interface{}) {
		fmt.Println(args...)
	}, []interface{}{"hello", "world", "test"})
	dm.AddTaskAfter(time.Second*30, "test4", func(args ...interface{}) {
		sum := 0
		for arg := range args {
			sum += arg
		}
		fmt.Println("sum : ", sum)
	}, []interface{}{1, 2, 3})

	time.AfterFunc(time.Second*40, func() {
		dm.Close()
	})

	dm.Start()
}
