//============================================================================
// cycle.go
// 	Author      : jtpeller
// 	Date		: December 06, 2021
// 	Description : Cycle sort implementation. Good when space is costly
//============================================================================

package gosort

func Cycle(a []int64) []int64 {
	// init
	n := len(a)

	// loop to find cycles
	for i, v := range a {
		// where to put it?
		pos := i
		for _, v2 := range a[i+1 : n] {
			if v2 < v {
				pos++
			}
		}

		// if it's already there, it's not a cycle
		if pos == i {
			continue
		}

		// if it is, put it there (or after duplicates)
		for v == a[pos] {
			pos++
		}
		a[pos], v = v, a[pos]

		// rotate the rest
		for pos != i {
			pos = i
			for _, v2 := range a[i+1 : n] {
				if v2 < v {
					pos++
				}
			}

			for v == a[pos] { // place it
				pos++
			}
			a[pos], v = v, a[pos]
		}
	}

	return a
}
