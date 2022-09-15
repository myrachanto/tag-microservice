package main

import (
	"log"

	"github.com/myrachanto/microservice/tag/src/routes"
)

func init() {
	log.SetPrefix("tag microservice ")
}
func main() {
	log.Println("Server started")
	routes.ApiSingle()
}
