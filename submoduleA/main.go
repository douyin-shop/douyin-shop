package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from submoduleA!")
	})
	log.Fatal(http.ListenAndServe(":201", nil))
}
