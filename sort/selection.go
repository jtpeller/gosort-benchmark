//============================================================================
// selection.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : selection sort implementation
//============================================================================

package gosort

func Selection(a []int64) []int64 {
	n := len(a)
	for i := 0; i < n; i++ {
		minidx := i
		for j := i; j < n; j++ {
			if a[j] < a[minidx] {
				minidx = j
			}
		}
		a[i], a[minidx] = a[minidx], a[i] // swap
	}
	return a
}
