//============================================================================
// stooge.go
// 	Author      : jtpeller
// 	Date		: December 06, 2021
// 	Description : Stooge sort implementation
//============================================================================

package gosort

func Stooge(a []int64) []int64 {
	stooge(a, 0, len(a)-1)
	return a
}

func stooge(a []int64, l, h int) {
	if l >= h {
		return
	}

	// first < last, swap
	if a[l] > a[h] {
		a[l], a[h] = a[h], a[l]
	}

	// if more than 2 elements in the array
	if h-l+1 > 2 {
		t := (h - l + 1) / 3
		stooge(a, l, h-t) // first 2/3
		stooge(a, l+t, h) // last
		stooge(a, l, h-t)
	}
}
