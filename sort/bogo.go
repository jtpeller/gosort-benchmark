//============================================================================
// bogo.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : the best sorting algorithm
//============================================================================

package gosort

import (
	utils "gosort/utils"
	"math/rand"
)

func Bogo(a []int64) []int64 {
	for !utils.IsSorted(a) {
		// shuffle array
		for i := len(a) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			if i != j {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}
