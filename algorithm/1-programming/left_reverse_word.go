package programming

/*
纯数组搬移数据
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。定义一个函数实现字符串左旋转操作的功能。
比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"

示例1：
输入：s = "abcdefg", k = 2
输出："cdefgab"

示例2：
输入：s = "lrloseumgh", k = 6
输出："umghlrlose"
*/

// 时间复杂度O(len)，空间复杂度O(len)，len表示字符串长度
func leftReverseWord(str string, num int) string {
	length := len(str)
	result := make([]byte, length, length)

	idx := num % length

	// idx - len-1
	for i := idx; i < length; i++ {
		result[i-idx] = str[i]
	}

	// 0 - idx
	for i := 0; i < idx; i++ {
		result[length-idx+i] = str[i]
	}

	return string(result)
}
