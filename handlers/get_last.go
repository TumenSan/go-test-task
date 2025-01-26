package handlers

import (
	"encoding/json"
	"net/http"

	"Go/database"
	"Go/models"
)

func GetLastHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	count := r.URL.Query().Get("count")
	if count == "" {
		count = "10"
	}

	rows, err := database.DB.Query("SELECT id, from_wallet, to_wallet, amount, time FROM transactions ORDER BY id DESC LIMIT $1", count)
	if err != nil {
		http.Error(w, "Failed to fetch transactions", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.From, &t.To, &t.Amount, &t.Time); err != nil {
			http.Error(w, "Failed to parse transaction", http.StatusInternalServerError)
			return
		}
		transactions = append(transactions, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}