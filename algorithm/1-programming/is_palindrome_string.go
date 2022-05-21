package programming

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写
说明：将空字符串定义为有效的回文串

示例1：
输入："A man, a plan, a canal: Panama"
输出：true

示例2：
输入："race a car"
输出：false
*/

// 时间复杂度O(n)，空间复杂度O(1)
func isPalindromeString(str string) bool {
	i, j := 0, len(str)-1
	for i < j {
		if !isAlpha(str[i]) {
			i++
			continue
		}
		if !isAlpha(str[j]) {
			j--
			continue
		}

		if toLower(str[i]) != toLower(str[j]) {
			return false
		}

		i++
		j--
	}
	return true
}

// 判断是不是数字或者字母
func isAlpha(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}
	if b >= 'A' && b <= 'Z' {
		return true
	}
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

// 大写转小写
func toLower(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b
	}
	if b >= '0' && b <= '9' {
		return b
	}
	return b + 32
}
