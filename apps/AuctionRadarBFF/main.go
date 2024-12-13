package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World!")
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running BFF....")
	http.ListenAndServe(":4000", nil)
}
