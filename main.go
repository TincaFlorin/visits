package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	counter := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>You have %d visits</h1>", counter)
		counter++
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
