//============================================================================
// tim.go
// 	Author      : jtpeller
// 	Date		: December 06, 2021
// 	Description : Tim sort implementation
//============================================================================

package gosort

import (
	"math"
)

const MIN int = 32

func TimSort(a []int64) []int64 {
	n := len(a)
	min := minrun(n)
	out := make([]int64, 0)

	for l := 0; l < n; l += min {
		r := int(math.Min(float64(l+min-1), float64(n-1)))
		out = timInsertion(a, l, r)
	}

	s := min
	for s < n {
		for l := 0; l < n; l += 2 * s {
			m := int(math.Min(float64(l+s-1), float64(n-1)))
			r := int(math.Min(float64(l+2*s-1), float64(n-1)))
			out = timMerge(a, l, m, r)
		}
		s *= 2
	}

	return out
}

func minrun(n int) int {
	r := 0
	for n >= MIN {
		r |= n & 1
		n >>= 1
	}
	return n + r
}

func timInsertion(a []int64, l, r int) []int64 {
	for i := l + 1; i < r+1; i++ {
		v := a[i]
		j := i - 1
		for j >= l && v < a[j] {
			a[j+1] = a[j]
			j -= 1
		}
		a[j+1] = v
	}
	return a
}

func timMerge(a []int64, l, m, r int) []int64 {
	// init
	n1 := m - l + 1
	n2 := r - m
	left := make([]int64, 0)
	right := make([]int64, 0)

	for i := 0; i < n1; i++ {
		left = append(left, a[l+i])
	}
	for i := 0; i < n2; i++ {
		right = append(right, a[m+1+i])
	}

	// loop to merge
	i, j, k := 0, 0, l

	for i < n1 && j < n2 {
		if left[i] < right[j] {
			a[k] = left[i]
			i++
		} else {
			a[k] = right[j]
			j++
		}
		k++
	}

	// copy remaining elements
	for i < n1 {
		a[k] = left[i]
		k++
		i++
	}

	for j < n2 {
		a[k] = right[j]
		k++
		j++
	}

	return a
}
