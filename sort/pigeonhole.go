//============================================================================
// pigeonhole.go
// 	Author      : jtpeller
// 	Date		: December 06, 2021
// 	Description : Pigeonhole sort implementation
//============================================================================

package gosort

import utils "gosort/utils"

func Pigeonhole(a []int64) []int64 {
	// init
	min := utils.Min(a...)
	max := utils.Max(a...)
	n := max - min + 1
	holes := make([]int64, n)

	// populate the holes
	for _, v := range a {
		holes[v - min]++
	}

	// put elements in order
	idx := 0
	for i := int64(0); i < n; i++ {
		for holes[i] > 0 {
			holes[i]--
			a[idx] = i + min
			idx++
		}
	}

	return a
}