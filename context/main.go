package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background() // root context
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel() // essential. internal timer was working for ctx and it should be turned off.

	sleepAndTalk(ctx, 5*time.Second, "hello")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}
