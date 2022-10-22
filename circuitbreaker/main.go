package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	sender := Breaker(send, 10)

	for i:=0;i<50;i++ {
		result, err := sender(ctx)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(result)
	}
}

type Circuit func(ctx context.Context) (string, error)

func Breaker(circuit Circuit, failureThreshold uint) Circuit {
	var m sync.RWMutex
	consecutiveFailures := 0
	lastAttempt := time.Now()

	return func(ctx context.Context) (string, error) {
		m.RLock()

		d := consecutiveFailures - int(failureThreshold)
		if d >= 0 {
			shouldRetryAt := lastAttempt.Add(time.Second * 2 << d)
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()

				return "", errors.New("service unreachable")
			}
		}

		m.RUnlock()

		response, err := circuit(ctx)

		m.Lock()
		defer m.Unlock()
		lastAttempt = time.Now()

		if err != nil {
			consecutiveFailures++
			return response, err
		}

		consecutiveFailures = 0

		return response, nil
	}
}

func send(ctx context.Context) (string, error) {
	return "", errors.New("bad request")
}