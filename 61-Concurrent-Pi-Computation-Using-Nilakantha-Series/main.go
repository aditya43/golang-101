// The number π is a mathematical constant. It is defined as the ratio of a circle's circumference to its diameter, and it also has various equivalent definitions. π is an irrational number. It's decimal representation never ends and it never settles into permanent repeating pattern. Here, we are using "Nilakantha Series" (which is over 500 years old) to compute π. "Nilakantha" was an astronomer and mathematician from Kerala school of astronomy and mathematics. "Nilakantha Series" is known to be a fastest algorithm to calculate value of π.

// A concurrent computation of pi using Nilakantha's formula.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// We want to create a virtual scoreboard where we can simutaneously
// show the current value of Pi and how many Nilakantha terms we
// have calculated.

// We can create the virtual scoreboard by using some ANSI escape codes
// Here's a wikipedia article giving you the complete list of ANSI escape
// codes: https://en.wikipedia.org/wiki/ANSI_escape_code
const ANSIClearScreenSequence = "\033[H\033[2J"
const ANSIFirstSlotScreenSequence = "\033[2;0H"
const ANSISecondSlotScreenSequence = "\033[3;0H"

// The channel used to update the current value of pi on the scoreboard
var pichan chan float64 = make(chan float64)

// The channel that we use to indicate that the program can exit
var computationDone chan bool = make(chan bool, 1)

// Number of Nilakantha terms for the scoreboard
var termsCount int

// This function serves as our virtual scoreboard to show the current
// computed value of Pi using Nilakantha's formula
func printCalculationSummary() {

	fmt.Print(ANSIClearScreenSequence)
	fmt.Println(ANSIFirstSlotScreenSequence, "Computed Value of Pi:\t\t", <-pichan)
	fmt.Println(ANSISecondSlotScreenSequence, "# of Nilakantha Terms:\t\t", termsCount)
}

// We are going to use Nilakantha's formula from the Tantrasamgraha (which
// is more than 500 years old) to compute the value of Pi to several decimal
// points
func pi(n int) float64 {
	ch := make(chan float64)
	// The k variable is considered to be the current step
	for k := 1; k <= n; k++ {
		// Each Nilakantha term is calculated in its own goroutine
		go nilakanthaTerm(ch, float64(k))
	}
	f := 3.0

	// We sum up the calculated Nilakantha terms for n steps
	for k := 1; k <= n; k++ {
		termsCount++
		f += <-ch
		pichan <- f
	}

	// We notify that the computation is done over the channel
	computationDone <- true
	return f
}

// This function gives us the nilakanthaTerm for the kth step
func nilakanthaTerm(ch chan float64, k float64) {
	j := 2 * k

	if int64(k)%2 == 1 {
		ch <- 4.0 / (j * (j + 1) * (j + 2))
	} else {
		ch <- -4.0 / (j * (j + 1) * (j + 2))
	}
}

func main() {

	// We use a ticker to specify the interval to update the values
	// on the scoreboard
	ticker := time.NewTicker(time.Millisecond * 108)

	// If the user wants to prematurely break out of the program by
	// issuing a Ctrl+C, we tell the signal package to notify us over
	// this interrupt channel
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go pi(5000)

	// This anonymous function is responsible for updating the scoreboard
	// as per the interval specified by the ticker
	go func() {
		for range ticker.C {
			printCalculationSummary()
		}
	}()

	for {
		select {
		// Handle the case when the computation has ended, we can
		// end the program (exit out of this infinite loop)
		case <-computationDone:
			ticker.Stop()
			fmt.Println("Program done calculating Pi.")
			os.Exit(0)

		// If the user interrupts the program (Ctrl+C) we end the
		// program (exit out of this infinite loop)
		case <-interrupt:
			ticker.Stop()
			fmt.Println("Program interrupted by the user.")
			return
		}
	}
}
