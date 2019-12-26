package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const usage = `usage: barman [-h] [COORDINATE [COORDINATE ...]]

Tiny terminal histogram creator.

positional arguments:
  COORDINATE            the x coordinate of the bar

optional arguments:
  -h                    show this help message and exit`

var bars = [8]rune{'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println()
		os.Exit(0)
	}

	if args[0] == "-h" {
		fmt.Println(usage)
		os.Exit(0)
	}

	numbers, err := parse(args)
	if err != nil {
		fmt.Println("invalid arguments")
		os.Exit(1)
	}

	chart := Plot(numbers)
	fmt.Println(chart)
}

func parse(args []string) ([]float64, error) {
	xs := make([]float64, 0)

	for _, arg := range args {
		x, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return nil, err
		}

		xs = append(xs, x)
	}

	return xs, nil
}

//Plot makes a bar plot.
func Plot(xs []float64) string {
	builder := new(strings.Builder)

	min, max := minmax(xs)

	if min == max {
		for range xs {
			builder.WriteRune(bars[len(bars)/2])
		}
		return builder.String()
	}

	for _, x := range xs {
		i := int((x - min) / (max - min) * float64(len(bars)-1))
		builder.WriteRune(bars[i])
	}

	return builder.String()
}

func minmax(numbers []float64) (float64, float64) {
	min, max := numbers[0], numbers[0]

	for _, number := range numbers {
		if max < number {
			max = number
		}
		if min > number {
			min = number
		}
	}

	return min, max
}
