package programming

/*
顺序扫描下标操作
给定一个有序数组 nums，请原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
不要使用额外的数组空间，必须在原地修改输入数组，并在使用O(1)额外空间的条件下完成。

示例1：
输入：nums = [1, 1, 2]
输出：2, nums = [1, 2]
*/

// 时间复杂度O(n)，空间复杂度O(1)
func removeDuplicates(nums []int) int {
	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}
