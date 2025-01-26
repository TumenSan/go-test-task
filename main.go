package main

import (
	"Go/database"
	"Go/handlers"
	"log"
	"net/http"
)

func main() {
	// Инициализация базы данных
	database.InitDatabase()

	// Регистрация HTTP-обработчиков
	http.HandleFunc("/api/wallet/", handlers.GetBalanceHandler)   // Обработчик для получения баланса кошелька
	http.HandleFunc("/api/send", handlers.SendHandler)            // Обработчик для выполнения перевода
	http.HandleFunc("/api/transactions", handlers.GetLastHandler) // Обработчик для получения последних транзакций

	// Запуск HTTP-сервера на порту 8080
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
