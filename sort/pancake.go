//============================================================================
// pancake.go
// 	Author      : jtpeller
// 	Date		: December 06, 2021
// 	Description : pancake sort implementation
//============================================================================

package gosort

import utils "gosort/utils"

func Pancake(a []int64) []int64 {
	s := len(a)

	for s > 1 {
		idx, _ := utils.MaxIdx(a[:s]...)

		if idx != s-1 { // if it is not at the end...
			flip(a, idx) // move max to the end requires moving max to the beginning
			flip(a, s-1) // move max to the end via reversal
		}
		s--
	}

	return a
}

func flip(a []int64, i int) {
	s := 0
	for s < i {
		a[s], a[i] = a[i], a[s]
		s++
		i--
	}
}
