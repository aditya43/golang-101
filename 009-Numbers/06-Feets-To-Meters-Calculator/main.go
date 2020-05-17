// OS Package: strconv allows us to convert vasic values from/to a string value.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	feets, _ := strconv.ParseFloat(arg, 64)

	meters := feets * 0.3048

	fmt.Printf("\n%g feets equals to %g meters\n", feets, meters)
}
