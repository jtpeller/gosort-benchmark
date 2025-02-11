//============================================================================
// heapsort.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : heap sort implementation (based on golang's heapSort in
// 				  sort.go, modified slightly)
//============================================================================

package gosort

import "math"

func HeapSort(a []int64) []int64 {
	heapify(a)

	end := len(a) - 1
	for end > 0 {
		a[end], a[0] = a[0], a[end]
		end--
		siftDown(a, 0, end)
	}
	return a
}

func parent(i int) int {
	return int(math.Floor(float64(i-1) / 2))
}

func child(i int) int {
	return 2*i + 1
}

// enforces heap property
func heapify(a []int64) {
	n := len(a)
	s := parent(n - 1)

	for s >= 0 {
		siftDown(a, s, n-1)
		s--
	}
}

// performs the work to actually enforce the heap property
func siftDown(a []int64, s, e int) {
	root := s
	for child(root) <= e {
		child := child(root)
		idx := root

		if a[idx] < a[child] {
			idx = child
		}

		if child+1 <= e && a[idx] < a[child+1] {
			idx = child + 1
		}

		if idx == root {
			return
		}

		a[root], a[idx] = a[idx], a[root]
		root = idx
	}
}
