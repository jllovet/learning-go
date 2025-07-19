package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	router.ServeFiles("/static/*filepath", http.Dir(dir))
	log.Fatal(http.ListenAndServe(":8000", router))
}
