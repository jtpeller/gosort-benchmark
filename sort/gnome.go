//============================================================================
// gnome.go
// 	Author      : jtpeller
// 	Date		: December 03, 2021
// 	Description : gnome sort implementation
//============================================================================

package gosort

func Gnome(a []int64) []int64 {
	n := len(a)
	idx := 0

	for idx < n {
		if idx == 0 {
			idx++
		}
		if a[idx] >= a[idx-1] {
			idx++
		} else {
			a[idx], a[idx-1] = a[idx-1], a[idx]
			idx--
		}
	}
	return a
}
