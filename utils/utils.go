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

func Max(nums ...int64) int64 {
	max := nums[0]
	for _, v := range nums {
		if max < v {
			max = v
		}
	}
	return max
}
