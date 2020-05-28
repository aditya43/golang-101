package main

import "fmt"

// Multiple concurrent workers pulling items off the queue
func main() {
	jobs := make(chan int, 100)    // Buffered channel
	results := make(chan int, 100) // Buffered channel

	go worker(jobs, results) // Deploying worker
	go worker(jobs, results) // Deploying worker
	go worker(jobs, results) // Deploying worker
	go worker(jobs, results) // Deploying worker
	go worker(jobs, results) // Deploying worker
	go worker(jobs, results) // Deploying worker
	go worker(jobs, results) // Deploying worker
	go worker(jobs, results) // Deploying worker

	for i := 0; i <= 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j <= 100; j++ {
		fmt.Println(<-results)
	}
}

// jobs: channel to take jobs to do
// results: channel to send results on
//
// Instead of specifying bi-directional channels,
// "jobs <-chan int": we will only every receive from the "jobs" channel.
// "results chan<- int": we we will only ever send on the results channel
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

// Calculate a fibonacci series for a given number
func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
