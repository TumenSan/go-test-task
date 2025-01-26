package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"Go/database"
)

func SendHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		From   string  `json:"from"`
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	tx, err := database.DB.Begin()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	var fromBalance float64
	if err := tx.QueryRow("SELECT balance FROM wallets WHERE address = $1", data.From).Scan(&fromBalance); err != nil {
		http.Error(w, "Sender wallet not found", http.StatusNotFound)
		return
	}

	if fromBalance < data.Amount {
		http.Error(w, "Insufficient funds", http.StatusBadRequest)
		return
	}

	if _, err := tx.Exec("UPDATE wallets SET balance = balance - $1 WHERE address = $2", data.Amount, data.From); err != nil {
		http.Error(w, "Failed to update sender balance", http.StatusInternalServerError)
		return
	}

	if _, err := tx.Exec("UPDATE wallets SET balance = balance + $1 WHERE address = $2", data.Amount, data.To); err != nil {
		http.Error(w, "Failed to update recipient balance", http.StatusInternalServerError)
		return
	}

	if _, err := tx.Exec("INSERT INTO transactions (from_wallet, to_wallet, amount, time) VALUES ($1, $2, $3, $4)", data.From, data.To, data.Amount, time.Now().Format(time.RFC3339)); err != nil {
		http.Error(w, "Failed to record transaction", http.StatusInternalServerError)
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)
}