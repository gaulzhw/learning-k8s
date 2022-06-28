package channel

import (
	"fmt"
	"time"
)

/*
有 4 个 goroutine，编号为 1、2、3、4。
每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写程序，让输出的编号总是按照 1、2、3、4、1、2、3、4……这个顺序打印出来。
*/

type Token struct{}

func newWorker(id int, ch chan Token, nextCh chan Token) {
	for {
		token := <-ch // 取得令牌
		fmt.Println(id + 1) // id从1开始
		time.Sleep(time.Second)
		nextCh <- token
	}
}
