//============================================================================
// bubble.cpp
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : Bubble sort implementation
//============================================================================

package gosort

func BubbleSort(nums []int64) []int64 {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n - i - 1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func BubbleSortFloat(nums []float64) []float64 {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n - i - 1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}