//============================================================================
// Merge.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : Merge sort implementation
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
		a[i], a[minidx] = a[minidx], a[i]	// swap
	}
	return a
}

func SelectionFloat(a []float64) []float64 {
	n := len(a)
	for i := 0; i < n; i++ {
		minidx := i
		for j := i; j < n; j++ {
			if a[j] < a[minidx] {
				minidx = j
			}
		}
		a[i], a[minidx] = a[minidx], a[i]	// swap
	}
	return a
}
