package programming

/*
字符串中元素替换，减少数组元素的搬移
实现一个函数，把字符串 s 中的每个空格替换成 "%20"。

限制：0 <= len(s) <= 10000

示例1：
输入：s = "We are happy."
输出："We%20are%20happy."
*/

// 时间复杂度O(n)，空间复杂度O(n)
func replaceSpace(str string) string {
	// 计算新字符串长度
	spaceCount := 0
	for _, b := range str {
		if b == ' ' {
			spaceCount++
		}
	}
	result := make([]byte, 0, len(str)+spaceCount*2)

	// 给字符串赋值
	for _, b := range str {
		if b != ' ' {
			result = append(result, byte(b))
		} else {
			result = append(result, '%')
			result = append(result, '2')
			result = append(result, '0')
		}
	}
	return string(result)
}
