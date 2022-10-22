package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()

	wrapped := DebounceFirst(send, time.Second*120)

	for i := 0; i < 50; i++ {
		result, err := wrapped(ctx)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}

	}

	time.Sleep(time.Second * 120)

	fmt.Println("Done")
}

type Circuit func(ctx context.Context) (string, error)

func DebounceFirst(circuit Circuit, d time.Duration) Circuit {
	var threshold time.Time
	var result string
	var err error
	var m sync.Mutex

	return func(ctx context.Context) (string, error) {
		m.Lock()
		defer func() {
			threshold = time.Now().Add(d)
			m.Unlock()
		}()

		if time.Now().Before(threshold) {
			return result, err
		}

		result, err = circuit(ctx)
		return result, err
	}
}

func send(ctx context.Context) (string, error) {
	return "Successful response", nil
}
