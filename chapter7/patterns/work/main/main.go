// This sample program demonstrates how to use the work package
// to use a pool of goroutines to get work done.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/nuinattapon/code/chapter7/patterns/work"
)

// names provides a set of names to display.
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter provides special support for printing names.
type namePrinter struct {
	name string
}

// Task implements the Worker interface.
func (m *namePrinter) Task() {
	start := time.Now()
	time.Sleep(time.Millisecond *
		time.Duration(rand.Int63n(100)))
	elapsed := time.Since(start)
	log.Printf("Task '%s' - Elapsed time: %d ms", m.name, elapsed.Milliseconds())
}

// main is the entry point for all Go programs.
func main() {
	start := time.Now()

	// Create a work pool with 4 goroutines.
	// Raspberry Pi 4 has 4 cores
	p := work.New(4)
	// Assuming that each task takes 100 ms.
	// So total task needs 100 x 5 (names) x 100 ms = 50 sec
	// 2 work pools will finish in 25 sec
	// 3 work pools will finish in 17 sec
	// 4 work pools will finish in 12.5 sec

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// Iterate over the slice of names.
		for _, name := range names {
			// Create a namePrinter and provide the
			// specific name.
			np := namePrinter{
				name: name,
			}

			go func() {
				// Submit the task to be worked on. When RunTask
				// returns we know it is being handled.
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// Shutdown the work pool and wait for all existing work
	// to be completed.
	p.Shutdown()
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %d ms", elapsed.Milliseconds())
}
