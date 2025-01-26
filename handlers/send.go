package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"Go/database"
)

// SendHandler обрабатывает запрос на перевод средств между кошельками
// Возвращает HTTP-статус 200 в случае успеха или соответствующий код ошибки
func SendHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Декодирование JSON из тела запроса
	var data struct {
		From   string  `json:"from"`   // Адрес кошелька отправителя
		To     string  `json:"to"`     // Адрес кошелька получателя
		Amount float64 `json:"amount"` // Сумма перевода
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Начало транзакции
	tx, err := database.DB.Begin()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback() // Откат транзакции в случае ошибки

	// Получение баланса отправителя
	var fromBalance float64
	if err := tx.QueryRow("SELECT balance FROM wallets WHERE address = $1", data.From).Scan(&fromBalance); err != nil {
		http.Error(w, "Sender wallet not found", http.StatusNotFound)
		return
	}

	// Проверка достаточности средств на балансе отправителя
	if fromBalance < data.Amount {
		http.Error(w, "Insufficient funds", http.StatusBadRequest)
		return
	}

	// Списание средств с кошелька отправителя
	if _, err := tx.Exec("UPDATE wallets SET balance = balance - $1 WHERE address = $2", data.Amount, data.From); err != nil {
		http.Error(w, "Failed to update sender balance", http.StatusInternalServerError)
		return
	}

	// Зачисление средств на кошелек получателя
	if _, err := tx.Exec("UPDATE wallets SET balance = balance + $1 WHERE address = $2", data.Amount, data.To); err != nil {
		http.Error(w, "Failed to update recipient balance", http.StatusInternalServerError)
		return
	}

	// Запись транзакции в историю
	if _, err := tx.Exec("INSERT INTO transactions (from_wallet, to_wallet, amount, time) VALUES ($1, $2, $3, $4)", data.From, data.To, data.Amount, time.Now().Format(time.RFC3339)); err != nil {
		http.Error(w, "Failed to record transaction", http.StatusInternalServerError)
		return
	}

	// Фиксация транзакции
	tx.Commit()
	w.WriteHeader(http.StatusOK) // Успешное завершение
}
