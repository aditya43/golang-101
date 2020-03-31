package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	const (
		feetsInMeters = 0.3048
		feetsInYards  = feetsInMeters / 0.9144
	)

	args := os.Args[1]

	feets, _ := strconv.ParseFloat(args, 64)

	meeters := feets * feetsInMeters
	yards := feets * feetsInYards
	roundedYards := math.Round(feets * feetsInYards)

	fmt.Printf("\n\n%g feets equals to %g meters\n", feets, meeters)
	fmt.Printf("%g feets equals to %g yards\n", feets, yards)
	fmt.Printf("%g feets equals to %.1f yards\n", feets, yards)
	fmt.Printf("%g feets equals to %g roundedYards\n\n", feets, roundedYards)
}
