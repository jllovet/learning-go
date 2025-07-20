package main

import (
	"log"

	"github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
