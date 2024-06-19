package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", allRequests)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
func allRequests(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
