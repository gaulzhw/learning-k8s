package channel

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateCases(t *testing.T) {
	var ch1 = make(chan int, 10)
	var ch2 = make(chan int, 10)

	// 创建SelectCase
	var cases = createCases(ch1, ch2)

	// 执行10次select
	for i := 0; i < 10; i++ {
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() {
			// recv case
			fmt.Println("recv:", cases[chosen].Dir, recv, ok)
		} else {
			// send case
			fmt.Println("send:", cases[chosen].Dir, ok)
		}
	}
}
