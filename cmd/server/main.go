package main

import (
	"log"
	"net/http"
	"targeting-engine/internal/delivery"
)

func main() {
	http.HandleFunc("/v1/delivery", delivery.DeliveryHandler)

	log.Println("Server is running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}