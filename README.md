# gosort

Various sorting algorithms written in go

## Contents

 - sort
    - bitonic.go -- implementation for bitonic sort
    - bogo.go -- implementation for bogo sort
    - bubble.go -- implementation for bubble sort
    - bucket.go -- implementation for bucket sort
    - cocktail.go -- implementation for cocktail sort
    - comb.go -- implementation for comb sort
    - counting.go -- implementation for counting sort
    - cycle.go -- implementation for cycle sort
    - gnome.go -- implementation for gnome sort
    - heapsort.go -- implementation for heap sort
    - insertion.go -- implementation for insertion sort
    - merge.go -- implementation for merge sort
    - oddeven.go -- implementation for odd-even sort
    - pancake.go -- implementation for pancake sort
    - pigeonhole.go -- implementation for pigeonhole sort
    - quicksort.go -- implementation for quicksort
    - radix.go -- implementation for radix sort
    - selection.go -- implementation for selection sort
    - shell.go -- implementation for shell sort
    - stooge.go -- implementation for stooge sort
    - tim.go -- implementation for tim sort
 - utils
    - utils.go -- general utils and other useful things
 - go.mod -- module file
 - input.txt -- sample text file with numbers 1 thru 16384 (shuffled)
 - main.go -- houses main
 - README.md -- the thing you're reading
 - test.txt -- smaller test file with numbers 1 thru 100 (shuffled)

## Usage

Open the root folder of this repo and run:
`go run main.go`

This will run the program on the default file `test.txt`

Run this command for help on the possible options:
`go run main.go -h`

Instead of outputting the sorted values, `main.go` compares the running times for each of the different algorithms.

To compile into a binary, use:
`go build`