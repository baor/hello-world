package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world is starting...")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world!")
	})
	http.ListenAndServe(":8080", nil)
}
