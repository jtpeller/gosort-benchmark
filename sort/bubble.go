//============================================================================
// bubble.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : Bubble sort implementation
//============================================================================

package gosort

func BubbleSort(a []int64) []int64 {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}
