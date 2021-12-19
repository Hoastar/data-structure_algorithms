package main

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		// 防止溢出取中间值写成如下语句。
		middle := left + ((right - left) / 2)
		if nums[middle] > target { // 此时target在左区间，所以 [left, middle-1]
			right = middle - 1
		} else if nums[middle] < target { // 此时target在右区间，所以 [middle+1, right]
			left = middle + 1
		} else { // nums[middle] == target
			return middle
		}
	}
	// 未找到目标值
	return -1
}

func searchTwo(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		middle := left + ((right - left) / 2)
		if nums[middle] > target {
			right = middle // target 在左区间，在[left, middle)中
		} else if nums[middle] < target {
			left = middle + 1 // target 在右区间，在[middle + 1, right)中
		} else { // nums[middle] == target
			return middle // 数组中找到目标值，直接返回下标
		}
	}
	return -1
}
