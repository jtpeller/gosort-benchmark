# gosort

## Overview

This repo contains various sorting algorithms written in go.

Running `main.go` will compare the execution times for each of the different sorting algorithms in the `sort` folder, instead of outputting the sorted values.
I implemented multithreading to improve the overall runtime, instead of running each sort sequentially.

Use the `-h` or `-help` flags to see all options.

To compile into a binary, use the following while in the directory with `main.go`:
`go build`

## Modes

There are a few different modes and options
To run them, open the root folder of this repo locally, in a terminal.
Then, follow the instructions below depending on which mode you want to execute.

Notes:

- Bad sorts (Bogo, Stooge) will default to certain values due to their slowness.
- Bitonic will round up to the nearest power of 2, due to its algorithmic nature. To do this, it pads the input array with 0s.
- Count mode cannot be used with input mode.
- Verbose option can be used with any of the modes

### Default Mode

Default mode executes on the default (random) values.
In your terminal, run:

`go run main.go`

This will run the program on the default file: `test.txt`, which has 100 values.

### Count Mode

You can input a number of values to test the sorting algorithms against.
In your terminal, run:

`go run main.go -count ####`

where #### is replaced by the number you want to enter. For instance:

`go run main.go -count 69420`

will execute each of the sorting algorithms on 69420 randomly generated digits varying in size.

### Input Mode

You can also input a file of your own custom values to test the sorting algorithms against.
In your terminal, run:

`go run main.go -input filename.txt`

where `filename.txt` is the path to the file you want to use. For instance:

`go run main.go -input input.txt`

will use the file `input.txt` (which is provided in this repo).

### Verbose Option

Verbose is an option to see extra printing or debug information.
Apply 0 for quiet, or 1 for verbose:

`go run main.go -verbose 1`

Verbose can be added to other modes (i.e. count+verbose mode).

## Contents

- sort
  - `bitonic.go` -- implementation for bitonic sort
  - `bogo.go` -- implementation for bogo sort
  - `bubble.go` -- implementation for bubble sort
  - `bucket.go` -- implementation for bucket sort
  - `cocktail.go` -- implementation for cocktail sort
  - `comb.go` -- implementation for comb sort
  - `counting.go` -- implementation for counting sort
  - `cycle.go` -- implementation for cycle sort
  - `gnome.go` -- implementation for gnome sort
  - `heapsort.go` -- implementation for heap sort
  - `insertion.go` -- implementation for insertion sort
  - `merge.go` -- implementation for merge sort
  - `oddeven.go` -- implementation for odd-even sort
  - `pancake.go` -- implementation for pancake sort
  - `pigeonhole.go` -- implementation for pigeonhole sort
  - `quicksort.go` -- implementation for quicksort
  - `radix.go` -- implementation for radix sort
  - `selection.go` -- implementation for selection sort
  - `shell.go` -- implementation for shell sort
  - `stooge.go` -- implementation for stooge sort
  - `tim.go` -- implementation for tim sort
- utils
  - `utils.go` -- general utils and other useful things
- `go.mod` -- module file
- `input.txt` -- sample text file with numbers 1 thru 16384 (shuffled)
- `main.go` -- houses main
- `README.md` -- the thing you're reading
- `test.txt` -- smaller test file with numbers 1 thru 100 (shuffled)
