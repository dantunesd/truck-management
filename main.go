package main

import (
	"log"
	"net/http"
	"truck-management/truck-management/infrastructure"
)

func main() {

	http.ListenAndServe(":3000", infrastructure.HandlerFactory())

	log.Println("server running at port :3000")
}
