//============================================================================
// quicksort.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : Quicksort sort implementation
//============================================================================

package gosort

import "math/rand"

func Quicksort(a []int64) []int64 {
	// base case
	if len(a) < 2 {
		return a
	}

	// init
	l, r := 0, len(a) - 1
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

func QuicksortFloat(a []float64) []float64 {
	// base case
	if len(a) < 2 {
		return a
	}

	// init
	l, r := 0, len(a) - 1
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
	QuicksortFloat(a[:l])
	QuicksortFloat(a[l+1:])

	return a
}
