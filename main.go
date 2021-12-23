package main

import (
	"errors"
	"flag"
	"fmt"
	sort "gosort/sort"
	utils "gosort/utils"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var VERBOSE bool = false	// for warning or other helpful output
const WARN_NUMS = 32768		// count value for when a warning will issue about too big an input size

func toIntSlice(input []string) []int64 {
	nums := make([]int64, len(input))
	for i, v := range input {
		temp, err := strconv.ParseInt(v, 10, 64)
		utils.HandleError(err)
		nums[i] = temp
	}
	return nums
}

func checkSorted(foo []int64, name string) {
	if VERBOSE {
		if !utils.IsSorted(foo) {
			if VERBOSE {
				fmt.Println(foo)
			}
			fmt.Println("\t\t\t", name, "is not sorted")
		}
	}
}

func performSort(f func([]int64) []int64, name string, nums []int64) {
	start := time.Now()
	foo := f(nums)
	duration := time.Since(start)
	fmt.Println(name, "took:  \t", duration, "on", len(nums), "values")
	checkSorted(foo, name)
}

func runSorts(vals []string) {
	// BITONIC SORT
	performSort(sort.Bitonic, "Bitonic sort", toIntSlice(vals))

	// BOGO SORT
	performSort(sort.Bogo, "Bogo sort", toIntSlice(utils.Generate(10)))

	// BUBBLE SORT
	performSort(sort.BubbleSort, "Bubble sort", toIntSlice(vals))

	// BUCKET SORT
	performSort(sort.Bucket, "Bucket sort", toIntSlice(vals))

	// COCKTAIL SORT
	performSort(sort.Cocktail, "Cocktail sort", toIntSlice(vals))

	// COMB SORT
	performSort(sort.CombSort, "Comb sort", toIntSlice(vals))

	// COUNTING SORT
	performSort(sort.Counting, "Counting sort", toIntSlice(vals))

	// CYCLE SORT
	performSort(sort.Cycle, "Cycle sort", toIntSlice(vals))

	// GNOME SORT
	performSort(sort.Gnome, "Gnome sort", toIntSlice(vals))

	// HEAP SORT
	performSort(sort.HeapSort, "Heap sort", toIntSlice(vals))

	// INSERTION SORT
	performSort(sort.Insertion, "Insertion sort", toIntSlice(vals))

	// MERGE SORT
	performSort(sort.MergeSort, "Merge sort", toIntSlice(vals))

	// ODD EVEN
	performSort(sort.OddEven, "Odd-even sort", toIntSlice(vals))

	// PANCAKE SORT
	performSort(sort.Pancake, "Pancake sort", toIntSlice(vals))

	// PIGEONHOLE SORT
	performSort(sort.Pigeonhole, "Pigeonhole sort", toIntSlice(vals))

	// QUICKSORT
	performSort(sort.Quicksort, "Quicksort", toIntSlice(vals))

	// RANDOMIZED QUICKSORT
	performSort(sort.RandomQuicksort, "RNG Quicksort", toIntSlice(vals))

	// RADIX
	performSort(sort.Radix, "Radix sort", toIntSlice(vals))

	// SELECTION SORT
	performSort(sort.Selection, "Selection sort", toIntSlice(vals))

	// SHELL SORT
	performSort(sort.ShellSort, "Shell sort", toIntSlice(vals))

	// STOOGE SORT
	performSort(sort.Stooge, "Stooge sort", toIntSlice(utils.Generate(128)))

	// TIM SORT
	performSort(sort.TimSort, "Tim sort", toIntSlice(vals))
}

func runWithCount(inputsize int64) {
	vals := utils.Generate(inputsize)
	runSorts(vals)
}

func runWithFile(filename string) {
	// load file
	raw, err := ioutil.ReadFile(filename)
	utils.HandleError(err)
	dat := string(raw)
	vals := strings.Split(dat, "\n")

	runSorts(vals)
}

func main() {
	// program setup (flags)
	filename := flag.String("input", "test.txt", "Input file for the sorting algorithms")
	inputsize := flag.Int64("count", 0, "Number of input to compute. Cannot be used with -input")
	verbose := flag.Bool("verbose", false, "Extra printing or debug information. 0 for quiet. 1 for verbose.")

	// parse!
	flag.Parse()

	// handle verbose-ness
	if *verbose {
		VERBOSE = true
	} else if !*verbose {
		VERBOSE = false
	} else {
		utils.PrintWarning("Warning: Verbose not in specified range. Defaulting to quiet mode.")
		VERBOSE = false
	}
	
	if *filename != "test.txt" && *inputsize != 0 {
		utils.HandleError(errors.New("cannot use the -input and -count flags together"))
	} else if *filename == "test.txt" && *inputsize == 0 {
		utils.PrintWarning("Warning: Using default test file: test.txt")
		runWithFile(*filename)	
	} else if *inputsize == 0 {
		utils.PrintInfo("Running with input file: " + *filename)
		runWithFile(*filename)
	} else if *inputsize != 0 {
		if *inputsize >= WARN_NUMS {
			utils.PrintWarning("Warning: The value you provided may take a while to compute...")
		}
		if VERBOSE {
			utils.PrintDebug("Note: This input size will not be used for bad sorts. Bad sorts include: Bogo, Stooge")
		}
		fmt.Println("Running with input size:", *inputsize)
		runWithCount(*inputsize)
	}	
}
