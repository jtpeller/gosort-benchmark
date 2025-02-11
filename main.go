//============================================================================
// main.go
// 	Author      : jtpeller
// 	Date		: December 06, 2021
// 	Description : Runs the entire sorting program.
//============================================================================

package main

import (
	"errors"
	"flag"
	"fmt"
	utils "gosort-benchmark/utils"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"
	"time"

	sortpkg "github.com/jtpeller/gosort"
)

type Sorter struct {
	f func([]int64) []int64 // sorting function
	s string                // name of Sort to output to console
	v []int64
}

type Result struct {
	s string        // name of sort
	t time.Duration // time it took to sort
	n int           // number of numbers that were sorted
}

var VERBOSE bool = false // for warning or other helpful output
const WARN_NUMS = 32768  // count value for when a warning will issue about too big an input size
var wg sync.WaitGroup
var results []Result // results array

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
	if !utils.IsSorted(foo) {
		if VERBOSE {
			fmt.Println(foo)
		}
		utils.PrintWarning("\t\t\t" + name + " is not sorted")
	} else if VERBOSE {
		utils.PrintInfo("\t\t\t" + name + " is sorted")
	}
}

func performSort(f func([]int64) []int64, name string, nums []int64, c chan Result) {
	defer wg.Done()

	start := time.Now()
	foo := f(nums)
	duration := time.Since(start)

	checkSorted(foo, name)

	c <- Result{s: name, t: duration, n: len(foo)}
}

func runSorts(sorts []Sorter) {
	for _, v := range sorts {
		c := make(chan Result, 1)
		wg.Add(1)
		go performSort(v.f, v.s, v.v, c)

		val := <-c
		results = append(results, val)
		close(c)
	}

	// wait for sorting to complete
	wg.Wait()

	sort.Slice(results, func(i, j int) bool {
		return results[i].t < results[j].t
	})

	utils.PrintInfo("Sorting Completed. Here are the results...")

	for _, r := range results {
		w := tabwriter.NewWriter(os.Stdout, 20, 1, 1, ' ', 0)
		fmt.Fprintln(w, r.s, "\t", r.t, "\t", r.n, " values")
		w.Flush()
	}

}

func runWithCount(inputsize int64) {
	start := time.Now()

	vals := utils.Generate(inputsize)

	// set up sort array to loop thru
	sorts := createSorter(vals, true)
	runSorts(sorts)

	dur := time.Since(start)
	utils.PrintInfo("Total Running Time", dur.String())
}

func runWithFile(filename string) {
	start := time.Now()

	// load file
	raw, err := os.ReadFile(filename)
	utils.HandleError(err)
	dat := string(raw)
	vals := strings.Split(dat, "\n")

	// set up sort array to loop thru
	sorts := createSorter(vals, true)

	runSorts(sorts)

	dur := time.Since(start)
	utils.PrintInfo("Total Running Time", dur.String())
}

func main() {
	// program setup (flags)
	filename := flag.String("input", "test.txt", "Input file for the sorting algorithms. Cannot be used with -count.")
	inputsize := flag.Int64("count", 0, "Number of input to compute. Cannot be used with -input.")
	verbose := flag.Bool("verbose", false, "Extra printing or debug information.")

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
		s := fmt.Sprintf("Running with input size: %d", *inputsize)
		utils.PrintInfo(s)
		runWithCount(*inputsize)
	}
}

func createSorter(vals []string, includeBadSorts bool) []Sorter {
	// create Sorter arr
	sorts := make([]Sorter, 0)

	sorts = append(sorts, Sorter{f: sortpkg.Bitonic, s: "Bitonic Sort", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.BubbleSort, s: "Bubble", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.Bucket, s: "Bucket", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.Cocktail, s: "Cocktail", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.CombSort, s: "CombSort", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.Counting, s: "Counting", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.Cycle, s: "Cycle", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.Gnome, s: "Gnome", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.HeapSort, s: "HeapSort", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.Insertion, s: "Insertion", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.MergeSort, s: "Merge", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.OddEven, s: "OddEven", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.Pancake, s: "Pancake", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.Pigeonhole, s: "Pigeonhole", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.Quicksort, s: "Quicksort", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.RandomQuicksort, s: "RNG Quicksort", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.Radix, s: "Radix", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.Selection, s: "Selection", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.ShellSort, s: "ShellSort", v: toIntSlice(vals)})
	sorts = append(sorts, Sorter{f: sortpkg.TimSort, s: "TimSort", v: toIntSlice(vals)})

	if includeBadSorts {
		sorts = append(sorts, Sorter{f: sortpkg.Bogo, s: "Bogo", v: toIntSlice(utils.Generate(10))})
		sorts = append(sorts, Sorter{f: sortpkg.Stooge, s: "Stooge", v: toIntSlice(utils.Generate(128))})
	}

	return sorts
}
