// This sample program demonstrates how to use a channel to
// monitor the amount of time the program is running and terminate
// the program if it runs too long.
package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/nuinattapon/code/chapter7/patterns/runner"
)

// timeout is the number of second the program has to finish.
const timeout = 4 * time.Second

// main is the entry point for the program.
func main() {
	rand.Seed(time.Now().UnixNano())

	log.Println("Starting work.")

	// Create a new timer value for this run.
	r := runner.New(timeout)

	// Add the tasks to be run.
	r.Add(createTask(), createTask(), createTask())

	// Run the tasks and handle the result.
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

// createTask returns an example task that sleeps for the specified
// number of seconds based on the id.
func createTask() func(int) {
	return func(id int) {
		duration := rand.Int63n(2000)
		log.Printf("Processor - Task #%d - It will take %d ms", id, duration)
		time.Sleep(time.Duration(duration) * time.Millisecond)
	}
}
