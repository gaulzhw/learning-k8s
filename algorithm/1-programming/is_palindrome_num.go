package programming

/*
给你一个整数 x，如果 x 是一个回文整数，返回 true；否则，返回 false。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如：121是回文，123不是回文。

示例1：
输入：x = 121
输出：true

示例2：
输入：x = -121
输出：false
解释：从左向右读，为-121，从右向左读，为121-。

示例3：
输入：x = 10
输出：false
解释：从右向左读，为01。

示例4：
输入：x = -101
输出：false
*/

// 时间复杂度O(1)，空间复杂度O(1)
func isPalindromeNum(num int) bool {
	if num < 0 {
		return false
	}

	// 将num转化成 整数数组
	// -2147483648 - 2147483647，最多10位数字
	digits := make([]int, 0, 10)
	for i := 0; num != 0; i++ {
		digits = append(digits, num%10)
		num = num / 10
	}

	// 判断回文串
	length := len(digits)
	for i := 0; i < length/2; i++ {
		if digits[i] != digits[length-i-1] {
			return false
		}
	}
	return true
}
