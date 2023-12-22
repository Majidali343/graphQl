package main

import (
	"fmt"
	"log"

	handlers "graphQl/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/jokes", handlers.GetJokesHandler)

	r.POST("/graphql", handlers.GraphqlHandler)

	port := 8080
	address := fmt.Sprintf(":%d", port)
	log.Printf("Server started on %s", address)

	if err := r.Run(address); err != nil {
		log.Fatal(err)
	}
}
