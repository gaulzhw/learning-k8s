package channel

func mapChan(in <-chan interface{}, fn func(interface{}) interface{}) <-chan interface{} {
	out := make(chan interface{}) // 创建一个输出chan
	if in == nil {
		close(out)
		return out
	}

	go func() {
		defer close(out)
		// 从输入chan读取数据，执行业务操作，也就是map操作
		for v := range in {
			out <- fn(v)
		}
	}()
	return out
}

func reduceChan(in <-chan interface{}, fn func(r, v interface{}) interface{}) interface{} {
	if in == nil {
		return nil
	}

	out := <-in // 先读取第一个元素
	// 实现reduce的主要逻辑
	for v := range in {
		out = fn(out, v)
	}
	return out
}
