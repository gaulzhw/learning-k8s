package programming

import (
	"strings"
)

/*
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
为简单起见，标点符号和普通字母一样处理。
例如输入字符串 "I am a student."，输出 "student. a am I"。

说明：
1. 无空格字符构成一个单词。
2. 输入字符串可以在前面或者后面包含多余的空格，但是翻转后的字符不能包含。
3. 如果两个单词间有多余的空格，将翻转后单词间的空格减少到只含一个。

示例1：
输入："the sky is blue"
输出："blue is sky the"

示例2：
输入：" hello world! "
输出："world! hello"
解释：输入字符串可以在前面或者后面包含多余的空格，但是翻转后的字符不能包含。

示例3：
输入："a good    example"
输出："example good a"
解释：如果两个单词间有多余的空格，将翻转后单词间的空格减少到只含一个。
*/

// 双指针
// 倒序遍历字符串 s ，记录单词左右索引边界 i , j
// 每确定一个单词的边界，则将其添加至单词列表 res
// 最终，将单词列表拼接为字符串，并返回即可

// 时间复杂度O(n)，空间复杂度O(n)
func reverseWords(str string) string {
	trimed := strings.TrimSpace(str)

	result := strings.Builder{}
	// i 表示单词首
	// j 表示单词尾
	i := len(trimed) - 1
	j := i
	for i >= 0 {
		for i >= 0 && trimed[i] != ' ' {
			i--
		}
		result.WriteString(trimed[i+1:j+1] + " ")
		for i >= 0 && trimed[i] == ' ' {
			i--
		}
		j = i
	}

	return strings.TrimSpace(result.String())
}
