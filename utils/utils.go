package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	Black  = "\u001b[30m"
	Red    = "\u001b[31m"
	Yellow = "\u001b[33m"
	Green  = "\u001b[32m"
	Blue   = "\u001b[34m"
	Reset  = "\u001b[0m"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// primarily for debugging
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

// for pretty printing for the user
func HandleError(e error) {
	if e != nil {
		PrintError(e.Error())
		os.Exit(1)
	}
}

func Generate(size int64) []string {
	list := make([]string, size)
	for i := int64(0); i < size; i++ {
		seed := rand.NewSource(time.Now().UnixNano())
		rng := rand.New(seed)
		list[i] = strconv.FormatInt(rng.Int63n(size), 10)
	}
	return list
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

func MaxIdx(nums ...int64) (int, int64) {
	idx := 0
	max := nums[0]
	for i, v := range nums {
		if max < v {
			max = v
			idx = i
		}
	}
	return idx, max
}

func PrintDebug(msg ...string) {
	fmt.Println(Blue, msg, Reset)
}

func PrintInfo(msg ...string) {
	fmt.Println(Green, msg, Reset)
}

func PrintWarning(msg ...string) {
	fmt.Println(Yellow, msg, Reset)
}

func PrintError(msg ...string) {
	fmt.Println(Red, msg, Reset)
}

func RemoveElement(a []int64, idx int) []int64 {
	return append(a[:idx], a[idx+1:]...)
}
