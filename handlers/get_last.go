package handlers

import (
	"encoding/json"
	"net/http"

	"Go/database"
	"Go/models"
)

// GetLastHandler обрабатывает запрос на получение последних транзакций
// Принимает параметр `count` из query-строки, который определяет количество возвращаемых транзакций
// По умолчанию возвращает 10 последних транзакций
// Возвращает JSON-массив транзакций или соответствующий код ошибки
func GetLastHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Получение параметра `count` из query-строки
	count := r.URL.Query().Get("count")
	if count == "" {
		count = "10" // Значение по умолчанию, если параметр не указан
	}

	// Выполнение SQL-запроса для получения последних транзакций
	rows, err := database.DB.Query("SELECT id, from_wallet, to_wallet, amount, time FROM transactions ORDER BY id DESC LIMIT $1", count)
	if err != nil {
		http.Error(w, "Failed to fetch transactions", http.StatusInternalServerError)
		return
	}
	defer rows.Close() // Закрытие rows после завершения

	// Сканирование результата запроса в структуру Transaction
	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.From, &t.To, &t.Amount, &t.Time); err != nil {
			http.Error(w, "Failed to parse transaction", http.StatusInternalServerError)
			return
		}
		transactions = append(transactions, t)
	}

	// Установка заголовка Content-Type и возврат JSON-ответа
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}
