package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"Go/database"
	"Go/models"
)

// GetBalanceHandler обрабатывает запрос на получение баланса кошелька
// Возвращает JSON с адресом кошелька и его балансом или соответствующий код ошибки
func GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Извлечение адреса кошелька из URL
	address := strings.TrimPrefix(r.URL.Path, "/api/wallet/")
	address = strings.Split(address, "/")[0] // Убираем возможные дополнительные пути

	// Выполнение SQL-запроса для получения баланса кошелька
	var balance float64
	if err := database.DB.QueryRow("SELECT balance FROM wallets WHERE address = $1", address).Scan(&balance); err != nil {
		http.Error(w, "Wallet not found", http.StatusNotFound)
		return
	}

	// Установка заголовка Content-Type и возврат JSON-ответа
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Wallet{Address: address, Balance: balance})
}
