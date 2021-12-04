//============================================================================
// cocktail.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : cocktail sort implementation
//============================================================================

package gosort

import utils "gosort/utils"

func Cocktail(a []int64) []int64 {
	i := 0
	for !utils.IsSorted(a) {
		for i = 0; i < len(a) - 2; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
		for ; i > 0; i-- {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	return a
}
