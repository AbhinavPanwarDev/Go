package main

import (
	"fmt"
	"time"
)

func main() {
	// Create an unbuffered channel
	ch := make(chan string)

	// Create a buffered channel with capacity 2
	bufferedCh := make(chan int, 2)

	// Start a goroutine to send data to the unbuffered channel
	go func() {
		fmt.Println("Sending message to channel...")
		ch <- "Hello from goroutine!" // This will block until someone receives
		fmt.Println("Message sent!")
	}()

	// Receive from the unbuffered channel
	msg := <-ch
	fmt.Println("Received:", msg)

	// Working with buffered channels
	fmt.Println("\nBuffered channel example:")
	bufferedCh <- 1 // Won't block because buffer has space
	bufferedCh <- 2 // Won't block because buffer has space

	// Reading from buffered channel
	fmt.Println(<-bufferedCh) // 1
	fmt.Println(<-bufferedCh) // 2

	// Channel direction example
	go sender(ch)
	go receiver(ch)

	// Wait a bit to see the results
	time.Sleep(time.Second)

	// Channel with select statement
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 500)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(time.Millisecond * 300)
		ch2 <- "Message from channel 2"
	}()

	// Select statement to handle multiple channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		}
	}
}

// Sender function - can only send to channel
func sender(ch chan<- string) {
	ch <- "Message from sender"
}

// Receiver function - can only receive from channel
func receiver(ch <-chan string) {
	msg := <-ch
	fmt.Println("Receiver got:", msg)
}
