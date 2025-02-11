//============================================================================
// insertion.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : Insertion sort implementation
//============================================================================

package gosort

func Insertion(a []int64) []int64 {
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
	return a
}
