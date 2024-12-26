package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()
	respch := make(chan any, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go fetchUserLikes(userName, respch, wg)
	go fetchUserMatches(userName, respch, wg)
	wg.Wait()
	close(respch)

	fmt.Println("It took:", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Second * 100)
	return "John Doe"
}

func fetchUserLikes(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 100)
	respch <- "42"
	wg.Done()
}

func fetchUserMatches(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 150)
	respch <- "ALICE"
	wg.Done()
}
