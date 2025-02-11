//============================================================================
// shell.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : shell sort implementation
//============================================================================

package gosort

func ShellSort(a []int64) []int64 {
	n := len(a)
	h := 1
	for h < n/3 { // knuth's formula
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				a[j], a[j-h] = a[j-h], a[j]
			}
		}
		h /= 3
	}
	return a
}
