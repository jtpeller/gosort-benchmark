//============================================================================
// bitonic.go
// 	Author      : jtpeller
// 	Date		: December 06, 2021
// 	Description : Bitonic sort implementation
//============================================================================

package gosort

import "math"

func Bitonic(a []int64) []int64 {
	// enforce that the input size is a power of 2
	l := len(a)
	if l&(l-1) != 0 { // len not a power of 2
		// pad a with zeros until next power of 2
		a = make([]int64, int(math.Pow(2, math.Ceil(math.Log2(float64(l))))))
	}

	return bitonic(a, 0, len(a), true)
}

// input size MUST be a power of 2
func bitonic(a []int64, low, count int, dir bool) []int64 {
	if count > 1 {
		k := count / 2
		bitonic(a, low, k, true)
		bitonic(a, low+k, k, false)
		bitonicMerge(a, low, count, dir)
	}
	return a
}

func bitonicMerge(a []int64, low, count int, dir bool) {
	if count > 1 {
		k := count / 2
		for i := low; i < low+k; i++ {
			bitonicSwap(a, i, i+k, dir)
		}
		bitonicMerge(a, low, k, dir)
		bitonicMerge(a, low+k, k, dir)
	}
}

func bitonicSwap(a []int64, i, j int, dir bool) {
	if (dir && a[i] > a[j]) || (!dir && a[i] < a[j]) {
		a[i], a[j] = a[j], a[i]
	}
}
