package main

import (
	"context"
	"fmt"
	mylog "github.com/yhmin84/go-study/context/log"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", mylog.Decorate(handler))
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, 42, int64(100))
	mylog.Println(ctx, "handler started\n")
	defer mylog.Println(ctx, "handler ended\n")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "hello")
	case <-ctx.Done():
		err := ctx.Err()
		mylog.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
