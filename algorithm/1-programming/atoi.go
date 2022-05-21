package programming

import (
	"math"
	"strings"
)

/*
首先，该函数会根据需要丢弃无用的开头空格字符，直到找到第一个非空格的字符为止。
寻找到的第一个非空字符为正或者负号时，则将该负号与后面尽可能多的连续数字组合起来，作为该整数的正负号；加入第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
注意：加入该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
在任何情况下，若函数不能进行有效的转换时，请返回0。

说明：
假设我们的环境只能存储32位大小的有符号整数，那么其数值范围为[-2^31, 2^31 - 1]，如果数值超过这个范围，请返回 INT_MAX(2^31 - 1) 或 INT_MIN(-2^31)

示例1：
输入："42"
输出：42

示例2：
输入："   -42"
输出：-42

示例3：
输入："4193 with words"
输出：4193

示例4：
输入："words and 987"
输出：0

示例5：
输入："-91283472332"
输出：-2147483648
*/

func atoi(str string) int {
	// 处理空
	if len(str) == 0 {
		return 0
	}

	// 处理空格
	trimed := strings.TrimSpace(str)
	if len(trimed) == 0 {
		return 0
	}

	sign := 1
	i := 0
	if trimed[i] == '-' {
		sign = -1
		i++
	} else if trimed[i] == '+' {
		sign = 1
		i++
	}

	// 处理数字
	result := 0
	for i < len(trimed) && trimed[i] >= '0' && trimed[i] <= '9' {
		digit := trimed[i] - '0'
		if result > math.MaxInt32 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		if result == math.MaxInt32 {
			if sign == 1 && digit > 7 {
				return math.MaxInt32
			}
			if sign == -1 && digit > 8 {
				return math.MinInt32
			}
		}
		result = result*10 + int(digit)
		i++
	}

	return sign * result
}
