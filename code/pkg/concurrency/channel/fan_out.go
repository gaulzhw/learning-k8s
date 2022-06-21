package channel

func fanOut(ch <-chan interface{}, out []chan interface{}, async bool) {
	go func() {
		defer func() {
			// 推出时关闭所有的输出chan
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()

		for v := range ch {
			v := v
			for i := 0; i < len(out); i++ {
				i := i
				if async {
					go func() {
						out[i] <- v
					}()
				} else {
					out[i] <- v
				}
			}
		}
	}()
}
