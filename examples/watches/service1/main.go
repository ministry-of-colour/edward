// A simple executable that stays runnning until an interrupt is received
// Based on: https://gobyexample.com/signals
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is service1 at: %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting to listen on port 8080")
	http.ListenAndServe(":8080", nil)
}
