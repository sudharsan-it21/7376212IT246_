package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "katrina karif!")

	})

	http.HandleFunc("/dep", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "IT!")

	})

	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GREEN!")

	})

	fmt.Printf("Server running (port=8086), route: http://localhost:8086/helloworld\n")
	if err := http.ListenAndServe(":8086", nil); err != nil {
		log.Fatal(err)
	}
}
