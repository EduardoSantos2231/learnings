package main

import (
	"fmt"
	"rate-limiter/bucket"
	"time"
)

func main() {
	ourBucket := bucket.NewBucket(25, time.Second*5)
	for i := range 30 {
		if ourBucket.Allow() {
			fmt.Printf("[%d] [V] passou\n", i)
		} else {
			fmt.Printf("[%d] [X] barrado\n", i)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
