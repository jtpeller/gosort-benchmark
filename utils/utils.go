package utils

type List interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func Equal(i, j interface{}) bool {
	return i == j
}

