package channel

import (
	"fmt"
	"time"
)

func sequencePrint() {
	size := 4
	signChans := make([]chan struct{}, size)
	for i := range signChans {
		signChans[i] = make(chan struct{})
	}

	for i := 1; i <= size; i++ {
		go func(id int, sign chan struct{}) {
			for {
				<-sign
				fmt.Println(id)
			}
		}(i, signChans[i-1])
	}

	// controller
	ticker := time.NewTicker(time.Second)
	id := 1
	for {
		<-ticker.C
		signChans[id-1] <- struct{}{}
		id++
		if id > size {
			id = 1
		}
	}
}
