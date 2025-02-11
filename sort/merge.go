//============================================================================
// merge.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : merge sort implementation
//============================================================================

package gosort

// merge sort for int64
func MergeSort(a []int64) []int64 {
	la := len(a)
	if la < 2 {
		return a
	}
	left := MergeSort(a[:la/2])  // up to middle
	right := MergeSort(a[la/2:]) // past middle
	return merge(left, right)    // combine
}

// merges slices l & r
func merge(l, r []int64) []int64 {
	b := make([]int64, 0)
	i, j := 0, 0
	ll := len(l)
	lr := len(r)

	// merge arrays
	for i < ll && j < lr {
		if l[i] < r[j] {
			b = append(b, l[i])
			i++
		} else {
			b = append(b, r[j])
			j++
		}
	}

	// copy remaining from left
	for ; i < ll; i++ {
		b = append(b, l[i])
	}

	// copy remaining from right
	for ; j < lr; j++ {
		b = append(b, r[j])
	}
	return b
}
