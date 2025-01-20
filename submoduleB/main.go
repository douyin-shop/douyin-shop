package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from submoduleB!")
	})
	log.Fatal(http.ListenAndServe(":202", nil))
}
