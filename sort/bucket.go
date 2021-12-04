//============================================================================
// bucket.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : bucket sort implementation
//============================================================================

package gosort

import (
	utils "gosort/utils"
)

func Bucket(a []int64) []int64 {
	// init
	bucketSize := 10
	min, max := utils.Min(a...), utils.Max(a...)
	nbuckets := int(max-min)/bucketSize + 1

	// create buckets
	buckets := make([][]int64, nbuckets)
	for i := 0; i < nbuckets; i++ {
		buckets[i] = make([]int64, 0)
	}

	for _, n := range a {
		idx := int(n-min)/bucketSize
		buckets[idx] = append(buckets[idx], n)
	}

	// create sorted arr
	out := make([]int64, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			Insertion(bucket)
			out = append(out, bucket...)
		}
	}

	return out
}
