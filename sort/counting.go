//============================================================================
// countingsort.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : counting sort implementation
//============================================================================

package gosort

import (
	utils "gosort/utils"
)

func Counting(a []int64) []int64 {
	// init
	max := utils.Max(a...)
	counts := make([]int64, max+1)

	// generate counts
	for _, v := range a {
		counts[v]++
	}

	// add prev to curr
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	// generate output
	out := make([]int64, len(a))
	for i := 0; i < len(a); i++ {
		v := a[i]			// value to add
		p := counts[v] - 1	// position

		out[p] = v			// add the value
		counts[v]--
	}
	return out
}
