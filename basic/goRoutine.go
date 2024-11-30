package main

import (
	"fmt"
	"sync"
	"time"
)

func counter(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when function completes

	for i := 1; i <= 3; i++ {
		fmt.Printf("Routine %d: Count %d\n", id, i)
		time.Sleep(time.Millisecond * 500) // Simulate some work
	}
}

func main() {
	var wg sync.WaitGroup

	// Launch 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)          // Increment wait group counter
		go counter(i, &wg) // Launch goroutine
	}

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("All done!")
}
