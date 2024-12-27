package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	userID := 10
	value, err := fetchUserData(ctx, userID)
	if err != nil {
		fmt.Println(err)
	}	
	fmt.Println(value, err)
	fmt.Println("It took:", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	respch := make(chan Response)
	go func() {
		value, err := fetchThirdPartyThingWhichCanTakeAWhile()
		respch <- Response{value: value, err: err}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetchUserData: context deadline exceeded")
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}

func fetchThirdPartyThingWhichCanTakeAWhile() (int, error) {
	time.Sleep(time.Millisecond * 1800)
	return 42, nil
}
