package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/splits", splitsHandler)
	http.HandleFunc("/", rootHandler)

	port := ":8080"
	fmt.Printf("Server is listening on %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
