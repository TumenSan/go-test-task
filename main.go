package main

import (
	"Go/database"
	"Go/handlers"
	"log"
	"net/http"
)

func main() {
	database.InitDatabase()

	http.HandleFunc("/api/wallet/", handlers.GetBalanceHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
