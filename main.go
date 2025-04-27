package main

import (
	"fmt"
	"github.com/anuragjha/bhayia/services"
	"log"
	"net/http"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	services.SetupRoutes()

	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
