package channel

import (
	"fmt"
	"testing"
)

func TestMapReduce(t *testing.T) {
	in := asStream(nil, 1, 2, 3)

	// map操作：*10
	mapFn := func(v interface{}) interface{} {
		return v.(int) * 10
	}

	// reduce操作：对map的结果进行累加
	reduceFn := func(r, v interface{}) interface{} {
		return r.(int) + v.(int)
	}

	sum := reduceChan(mapChan(in, mapFn), reduceFn) // 返回累加结果
	fmt.Println(sum)
}
