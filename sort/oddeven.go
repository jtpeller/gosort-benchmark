//============================================================================
// oddeven.go
// 	Author      : jtpeller
// 	Date		: December 06, 2021
// 	Description : Odd-Even sort implementation
//============================================================================

package gosort

func OddEven(a []int64) []int64 {
	sorted := false
	n := len(a)

	for !sorted {
		sorted = true

		// odd
		for i := 1; i < n-1; i += 2 {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				sorted = false
			}
		}

		// even
		for i := 0; i < n-1; i += 2 {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	return a
}
