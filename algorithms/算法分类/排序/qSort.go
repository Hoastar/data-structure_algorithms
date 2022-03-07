package main

import (
	"fmt"
)

func main() {
	var Array = []int{3, 2, 1, 5, 6, 4}
	fmt.Println(Array)
	Array = sortArray(Array)
	fmt.Println(Array)
}

/*
快速排序，基于比较，不稳定算法，时间平均O(nlogn)，最坏O(n^2)，空间O(logn)
分治思想，选主元，依次将剩余元素的小于主元放其左侧，大的放右侧
然后取主元的前半部分和后半部分进行同样处理，直至各子序列剩余一个元素结束，排序完成

注意：
	小规模数据(n<100)，由于快排用到递归，性能不如插排
	进行排序时，可定义阈值，小规模数据用插排，往后用快排
	(小数，主元，大数)
*/
func sortArray(nums []int) []int {
	var quick func(nums []int, left, right int) []int
	quick = func(nums []int, left, right int) []int {
		// 递归终止条件
		if left > right {
			return nil
		}
		// 左右指针及主元
		i, j, pivot := left, right, nums[left]
		for i < j {
			// 寻找小于主元的右边元素
			for i < j && nums[j] >= pivot {
				j--
			}
			// 寻找大于主元的左边元素
			for i < j && nums[i] <= pivot {
				i++
			}
			// 交换i/j下标元素
			nums[i], nums[j] = nums[j], nums[i]
		}
		// 交换元素
		nums[i], nums[left] = nums[left], nums[i]
		quick(nums, left, i-1)
		quick(nums, i+1, right)
		return nums
	}
	return quick(nums, 0, len(nums)-1)
}
