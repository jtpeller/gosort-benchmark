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

	// INSERTION SORT
	start = time.Now()
	sort.Insertion(nums)
	duration = time.Since(start)
	fmt.Println("Insertion sort took: ", duration)
	
	// MERGE SORT
	start = time.Now()
	sort.MergeSort(nums)
	duration = time.Since(start)
	fmt.Println("Merge sort took: ", duration)

	// QUICKSORT
	start = time.Now()
	sort.Quicksort(nums)
	duration = time.Since(start)
	fmt.Println("Quicksort took: ", duration)

	// SELECTION SORT
	start = time.Now()
	sort.Selection(nums)
	duration = time.Since(start)
	fmt.Println("Selection took: ", duration)
}
