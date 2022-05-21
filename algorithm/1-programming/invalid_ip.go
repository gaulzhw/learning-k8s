package programming

/*
给你一个有效的IPv4地址 address，返回这个 IP 地址的无效化版本。
所谓无效化 IP 地址，其实就是用"[.]"代替每个"."。

示例1：
输入：address = "1.1.1.1"
输出："1[.]1[.]1[.]1"

示例2：
输入：address = "255.100.50.0"
输出："255[.]100[.]50[.]0
*/

// 时间复杂度O(n)
func invalidIP(address []byte) []byte {
	// . -> [.]，多了2个
	result := make([]byte, 0, len(address)+2*3)
	for _, add := range address {
		if add != '.' {
			result = append(result, add)
		} else {
			result = append(result, '[')
			result = append(result, add)
			result = append(result, ']')
		}
	}
	return result
}
