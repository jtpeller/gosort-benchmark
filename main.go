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

func getInput(input []string) []int64 {
	nums := make([]int64, len(input))
	for i, v := range input {
		temp, err := strconv.ParseInt(v, 10, 64)
		utils.CheckError(err)
		nums[i] = temp
	}
	return nums
}

func main() {
	raw, err := ioutil.ReadFile("input.txt")
	utils.CheckError(err)
	dat := string(raw)
	vals := strings.Split(dat, "\n")

	// BOGO SORT
	bogo := []int64{5, 2, 6, 4, 1, 7, 8, 9, 3}
	start := time.Now()
	sort.Bogo(bogo)
	duration := time.Since(start)
	fmt.Println("Bogo sort took: ", duration, "on 10 values")

	// BUBBLE SORT
	nums := getInput(vals)
	start = time.Now()
	sort.BubbleSort(nums)
	duration = time.Since(start)
	fmt.Println("Bubble sort took: ", duration)

	// BUCKET SORT
	nums = getInput(vals)
	start = time.Now()
	sort.Bucket(nums)
	duration = time.Since(start)
	fmt.Println("Bucket sort took: ", duration)

	// COCKTAIL SORT
	nums = getInput(vals)
	start = time.Now()
	sort.Cocktail(nums)
	duration = time.Since(start)
	fmt.Println("Cocktail sort took: ", duration)

	// COUNTING SORT
	nums = getInput(vals)
	start = time.Now()
	sort.Counting(nums)
	duration = time.Since(start)
	fmt.Println("Counting sort took: ", duration)

	// GNOME SORT
	nums = getInput(vals)
	start = time.Now()
	sort.Gnome(nums)
	duration = time.Since(start)
	fmt.Println("Gnome sort took: ", duration)

	// HEAP SORT
	nums = getInput(vals)
	start = time.Now()
	sort.HeapSort(nums, 0, len(nums))
	duration = time.Since(start)
	fmt.Println("Heap sort took: ", duration)

	// INSERTION SORT
	nums = getInput(vals)
	start = time.Now()
	sort.Insertion(nums)
	duration = time.Since(start)
	fmt.Println("Insertion sort took: ", duration)
	
	// MERGE SORT
	nums = getInput(vals)
	start = time.Now()
	sort.MergeSort(nums)
	duration = time.Since(start)
	fmt.Println("Merge sort took: ", duration)

	// QUICKSORT
	nums = getInput(vals)
	start = time.Now()
	sort.Quicksort(nums)
	duration = time.Since(start)
	fmt.Println("Quicksort took: ", duration)

	// RANDOMIZED QUICKSORT
	nums = getInput(vals)
	start = time.Now()
	sort.RandomQuicksort(nums)
	duration = time.Since(start)
	fmt.Println("Quicksort (random) took: ", duration)

	// RADIX
	nums = getInput(vals)
	start = time.Now()
	sort.Radix(nums)
	duration = time.Since(start)
	fmt.Println("Radix took: ", duration)

	// SELECTION SORT
	nums = getInput(vals)
	start = time.Now()
	sort.Selection(nums)
	duration = time.Since(start)
	fmt.Println("Selection took: ", duration)

	// SHELL SORT
	nums = getInput(vals)
	start = time.Now()
	sort.ShellSort(nums)
	duration = time.Since(start)
	fmt.Println("Shell sort took: ", duration)
}
