package api

import (
	"fmt"
	"net/http"
)

func GetTruckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getting a truck")
	}
}
