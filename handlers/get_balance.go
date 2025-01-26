package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"Go/database"
	"Go/models"
)

func GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Извлекаем адрес из URL
	address := strings.TrimPrefix(r.URL.Path, "/api/wallet/")
	address = strings.Split(address, "/")[0]

	var balance float64
	if err := database.DB.QueryRow("SELECT balance FROM wallets WHERE address = $1", address).Scan(&balance); err != nil {
		http.Error(w, "Wallet not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Wallet{Address: address, Balance: balance})
}