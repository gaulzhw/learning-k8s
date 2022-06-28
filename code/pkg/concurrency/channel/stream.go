package channel

func asStream(done <-chan struct{}, values ...interface{}) <-chan interface{} {
	s := make(chan interface{})
	go func() {
		defer close(s)
		for _, v := range values {
			select {
			case <-done:
				return
			case s <- v: // 将数组元素塞入到chan中
			}
		}
	}()
	return s
}

/*
1. takeN：只取流中的前 n 个数据
2. takeFn：筛选流中的数据，只保留满足条件的数据
3. takeWhile：只取前面满足条件的数据，一旦不满足条件，就不再取
4. skipN：跳过流中前几个数据
5. skipFn：跳过满足条件的数据
6. skipWhile：跳过前面满足条件的数据，一旦不满足条件，当前这个元素和以后的元素都会输出给 Channel 的 receiver
*/

func takeN(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{}) // 创建输出流
	go func() {
		defer close(takeStream)
		// 只读取前num个元素
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream: // 从输入流中读取元素
			}
		}
	}()
	return takeStream
}
