package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	handleChannel chan Splits = make(chan Splits)
	readyChannel  chan bool   = make(chan bool)
)

func main() {
	defer close(handleChannel)

	http.Handle("/", http.TimeoutHandler(http.HandlerFunc(rootHandler), 10*time.Second, "Request timed out"))
	http.Handle("/upload", http.TimeoutHandler(http.HandlerFunc(uploadHandler), 10*time.Second, "Request timed out"))
	http.Handle("/splits", http.TimeoutHandler(http.HandlerFunc(splitsHandler), 10*time.Second, "Request timed out"))

	port := ":8080"
	fmt.Printf("Server is listening on %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
