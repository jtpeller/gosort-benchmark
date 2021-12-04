//============================================================================
// radix.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : radix sort implementation
//============================================================================

package gosort

import (
	utils "gosort/utils"
)

// base 10 radix sort
func Radix(a []int64) []int64 {
	max := utils.Max(a...)
	n := len(a)
	sigdig := int64(1)
	temp := make([]int64, n)

	// loop to get to the largest sigdig
	for max / sigdig > 0 {
		// generate buckets
		bucket := make([]int64, 10)
		for i := 0; i < n; i++ {
			bucket[(a[i] / sigdig) % 10]++
		}

		// add prev to curr
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}

		// fill a temp array
		for i := n - 1; i >= 0; i-- {
			bucket[(a[i] / sigdig) % 10]--
			temp[bucket[(a[i] / sigdig) % 10]] = a[i]
		}

		for i := 0; i < n; i++ {
			a[i] = temp[i]
		}
		sigdig *= 10
	}
	return a
}
