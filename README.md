# gosort

Various sorting algorithms written in go

## Contents

 - sort
    - bubble.go -- implementation for bubble sort
    - counting.go -- implementation for counting sort
    - heapsort.go -- implementation for heap sort
    - insertion.go -- implementation for insertion sort
    - merge.go -- implementation for merge sort
    - quicksort.go -- implementation for quicksort
    - radix.go -- implementation for radix sort
    - selection.go -- implementation for selection sort
 - utils
    - utils.go -- general utils and other useful things
 - go.mod -- module file
 - input.txt -- sample text file with numbers 1 thru 10000
 - main.go -- houses main
 - README.md -- the thing you're reading

## Usage

Open the root folder of this repo and run:
`go run main.go`
Instead of outputting the sorted values, main.go compares the running times for each of the different algorithms.