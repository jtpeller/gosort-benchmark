//============================================================================
// comb.go
// 	Author      : jtpeller
// 	Date		: December 06, 2021
// 	Description : Comb sort implementation
//============================================================================

package gosort

func CombSort(a []int64) []int64 {
	// init
	n := len(a)
	gap := n
	swapped := true

	// loop
	for gap != 1 || swapped {
		gap = calcGap(gap)
		swapped = false

		// compare elements using current gap
		for i := 0; i < n-gap; i++ {
			if a[i] > a[i+gap] {
				a[i], a[i+gap] = a[i+gap], a[i]
				swapped = true
			}
		}
	}

	return a
}

func calcGap(gap int) int {
	gap = (gap * 10) / 13
	if gap < 1 {
		return 1
	}
	return gap
}
