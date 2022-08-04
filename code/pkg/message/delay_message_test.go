package message

import (
	"fmt"
	"testing"
	"time"
)

func TestDelayMessage(t *testing.T) {
	dm := NewDelayMessage()

	dm.AddTaskAfter(time.Second*1, "test1", func(args []any) {
		fmt.Println(args)
	}, []any{1, 2, 3})
	dm.AddTaskAfter(time.Second*1, "test2", func(args []any) {
		fmt.Println(args)
	}, []any{4, 5, 6})
	dm.AddTaskAfter(time.Second*2, "test3", func(args []any) {
		fmt.Println(args)
	}, []any{"hello", "world", "test"})
	dm.AddTaskAfter(time.Second*3, "test4", func(args []any) {
		sum := 0
		for _, arg := range args {
			sum += arg.(int)
		}
		fmt.Println("sum : ", sum)
	}, []any{1, 2, 3})

	time.AfterFunc(time.Second*5, func() {
		dm.Close()
	})

	dm.Start()
}
