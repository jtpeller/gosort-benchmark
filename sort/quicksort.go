//============================================================================
// quicksort.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : Quicksort sort implementation
//============================================================================

package gosort

import "math/rand"

func Quicksort(a []int64) []int64 {
	return quick(a, 0, len(a)-1)
}

func RandomQuicksort(a []int64) []int64 {
	// base case
	if len(a) < 2 {
		return a
	}

	// init
	l, r := 0, len(a)-1
	pivot := rand.Int() % len(a)
	a[pivot], a[r] = a[r], a[pivot]

	// loop & swap
	for i := range a {
		if a[i] < a[r] {
			a[l], a[i] = a[i], a[l]
			l++
		}
	}

	// 'final' swap
	a[l], a[r] = a[r], a[l]

	// recurrence
	Quicksort(a[:l])
	Quicksort(a[l+1:])

	return a
}

func quick(a []int64, l, r int) []int64 {
	if l < r {
		pivot := partition(a, l, r)
		a = quick(a, l, pivot-1)
		a = quick(a, pivot+1, r)
	}
	return a
}

func partition(a []int64, l, r int) int {
	for l < r {
		for l < r && a[l] <= a[r] {
			r--
		}
		if l < r {
			a[l], a[r] = a[r], a[l]
		}
		for l < r && a[l] <= a[r] {
			l++
		}
		if l < r {
			a[l], a[r] = a[r], a[l]
			r--
		}
	}
	return l
}
