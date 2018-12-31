# My Solutions to the Advent of Code

About the [Advent of Code](https://adventofcode.com):

> Advent of Code is an Advent calendar of small programming puzzles for a
> variety of skill sets and skill levels that can be solved in any
> programming language you like.

These are my solutions (written in Go) for some of these puzzles.

## How Is This Repo Organized

Each puzzle is inside its own folder following the convention `yyyy/dd`,
where `yyyy` is the year (i.e. `2018`), and `dd` is the puzzle day (i.e. `01`).

To show the solution for each puzzle, run:

```sh
# Go to the puzzle directory
$ cd 2018/01

# Runs the puzzle with the provided input file
$ go run main.go

# Runs the puzzle with a custom input file
$ go run main.go -input=./path/to/input/file

# Some puzzles might expose additional flags, see which flags by running
$ go run main.go -h
```

To run the tests for each puzzle:

```sh
$ go test
```

## License

Copyright (C) Daniel Fernandes Martins

Distributed under the New BSD License. See LICENSE for further details.
The puzzles themselves are designed by [Eric Wastl](https://twitter.com/ericwastl).
