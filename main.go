package main

import (
	"fmt"
	sort "gosort/sort"
	utils "gosort/utils"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func getInput(filename string) []int64 {
	raw, err := ioutil.ReadFile(filename)
	utils.CheckError(err)
	dat := string(raw)
	vals := strings.Split(dat, "\n")
	nums := make([]int64, len(vals))
	for i, v := range vals {
		temp, err := strconv.ParseInt(v, 10, 64)
		utils.CheckError(err)
		nums[i] = temp
	}
	return nums
}

func main() {
	nums := getInput("input.txt")

	// BUBBLE SORT
	start := time.Now()
	sort.BubbleSort(nums)
	duration := time.Since(start)
	fmt.Println("Bubble sort took: ", duration)

	// 
}