//============================================================================
// heapsort.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : heap sort implementation (based on golang's heapSort in
// 				  sort.go, modified slightly)
//============================================================================

package gosort

func HeapSort(a []int64, y, z int) []int64 {
	first := y
	low := 0
	high := z - y

	// build heap
	for i := (high - 1)/2; i >= 0; i-- {
		a = heapify(a, i, high, first)
	}

	// pop into end of data
	for i := high - 1; i >= 0; i-- {
		a[first], a[first+i] = a[first+i], a[first]
		a = heapify(a, low, i, first)
	}

	return a
}

// enforces heap property
func heapify(a []int64, low, high, first int) []int64 {
	root := low
	for {
		child := 2 * root + 1
		if child >= high {
			break
		} else if child+1 < high && a[first+child] < a[first+child+1] {
			child++
		} else if !(a[first+root] < a[first+child]) {
			return a
		}
		a[first+root], a[first+child] = a[first+child], a[first+root]
		root = child
	}
	return a
}
