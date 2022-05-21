package programming

/*
两数之和
给定一个整数数组 nums 和一个整数目标值 target，找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
假设每种输入中只回对应一个答案。但是数组中同一个元素在答案理不能重复出现。
可以按任意顺序返回答案。

示例1：
输入：nums = [2, 7, 11, 15], target = 9
输出：[0, 1]
解释：因为 nums[0] + nums[1] == 9, 返回 [0, 1]
*/

// use hashmap，暴力求解时间复杂度为O(n^2)，借助HashMap时间复杂度为O(n)
func sumOfTwo(ints []int, target int) []int {
	result := make([]int, 2)
	// key: 整数的值, value: 整数在数组的索引位置
	m := make(map[int]int)
	for idx, val := range ints {
		v, ok := m[target-val]
		if !ok {
			// 如果不存在，保存起来
			m[val] = idx
			continue
		}
		// 如果存在，保存结果
		result[0] = idx
		result[1] = v
		break
	}
	return nil
}
