# gosort

## Overview

This repo contains various sorting algorithms written in go.

Running `main.go` will compare the execution times for each of the different sorting algorithms in the `sort` folder, instead of outputting the sorted values.
I implemented multithreading to improve the overall runtime, instead of running each sort sequentially.

To compile into a binary, use the following while in the directory with `main.go`:
`go build`

## Modes

There are a few different modes and options. You can use the `-h` or `-help` flags to see all options.

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

```sh
go run main.go
```

This will run the program on the default file: `test.txt`, which has 100 values.

### Count Mode

You can input a number of values to test the sorting algorithms against.
In your terminal, run:

```sh
go run main.go -count num
```

where #### is replaced by the number you want to enter. For instance:

```sh
go run main.go -count 69420
```

will execute each of the sorting algorithms on 69420 randomly generated digits varying in size.

### Input Mode

You can also input a file of your own custom values to test the sorting algorithms against.
In your terminal, run:

```sh
go run main.go -input filename
```

where `filename` is the path to the file you want to use. For instance:

```sh
go run main.go -input input.txt
```

will use the file `input.txt` (which is provided in this repo).

### Verbose Option

Verbose is an option to see extra printing or debug information.
Apply the verbose option for additional information.

```sh
go run main.go -verbose
```

Verbose can be added to other modes (i.e., count + verbose mode).

## Contents

- sort -- folder containing the implementation for the sort in the filename
  - `bitonic.go`
  - `bogo.go`
  - `bubble.go`
  - `bucket.go`
  - `cocktail.go`
  - `comb.go`
  - `counting.go`
  - `cycle.go`
  - `gnome.go`
  - `heapsort.go`
  - `insertion.go`
  - `merge.go`
  - `oddeven.go`
  - `pancake.go`
  - `pigeonhole.go`
  - `quicksort.go`
  - `radix.go`
  - `selection.go`
  - `shell.go`
  - `stooge.go`
  - `tim.go`
- utils
  - `utils.go` -- general utils and other useful things
- `go.mod` -- Go module file
- `input.txt` -- sample text file with numbers 1 thru 16384 (shuffled)
- `main.go` -- implementation for CLI options, timing, etc.
- `README.md` -- what you're reading
- `test.txt` -- smaller test file with numbers 1 thru 100 (shuffled)
