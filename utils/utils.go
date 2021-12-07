package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	black = "\u001b[30m"
	red = "\u001b[31m"
	yellow = "\u001b[33m"
	green = "\u001b[32m"
	blue = "\u001b[34m"
	reset = "\u001b[0m"
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
		rand.Seed(time.Now().UnixNano())
		list[i] = strconv.FormatInt(rand.Int63n(size), 10)
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

func PrintDebug(msg string) {
	fmt.Println(blue + msg + reset)
}

func PrintInfo(msg string) {
	fmt.Println(green + msg + reset)
}

func PrintWarning(msg string) {
	fmt.Println(yellow + msg + reset)
}

func PrintError(msg string) {
	fmt.Println(red + msg + reset)
}

func RemoveElement(a []int64, idx int) []int64 {
    return append(a[:idx], a[idx+1:]...)
}
