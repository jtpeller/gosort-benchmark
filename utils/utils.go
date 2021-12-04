package utils

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func IsSorted(a []int64) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func Min(nums ...int64) int64 {
	min := nums[0]
	for _, v := range nums {
		if min > v {
			min = v
		}
	}
	return min
}

func Max(nums ...int64) int64 {
	max := nums[0]
	for _, v := range nums {
		if max < v {
			max = v
		}
	}
	return max
}
