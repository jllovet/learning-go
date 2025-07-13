package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/randint", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Intn(100))
	})
	mux.HandleFunc("/randfloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})
	http.ListenAndServe(":8000", mux)
}
