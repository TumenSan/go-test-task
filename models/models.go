package models

type Wallet struct {
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
}

type Transaction struct {
	ID     int     `json:"id"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
	Time   string  `json:"time"`
}
