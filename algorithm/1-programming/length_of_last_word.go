package programming

import (
	"strings"
)

/*
给定一个字符串 s，由若干单词组成，单词之间用空格隔开。返回字符串中最后一个单词的长度。如果不存在最后一个单词，返回 0。
单词：指仅由字母组成、不包含任何空格字符的最大子字符串。

提示：
1. 1 <= s.length <= 10^4
2. s 仅有英文字母和空格组成

示例1：
输入：s = "Hello World"
输出：5

示例2：
输入：s = " "
输出：0
*/

// 时间复杂度O(n)，空间复杂度O(1)
func lengthOfLastWord(str string) int {
	trimed := strings.TrimSpace(str)

	i := len(trimed) - 1
	for i >= 0 && str[i] == ' ' {
		i--
	}

	if i < 0 {
		return 0
	}

	result := 0
	for i >= 0 && trimed[i] != ' ' {
		result++
		i--
	}
	return result
}
