package main

import (
	"log"
	"truck-management/truck-management/infrastructure"
)

func main() {
	if err := infrastructure.InitializeWebServer(); err != nil {
		log.Fatal(err)
	}
}
