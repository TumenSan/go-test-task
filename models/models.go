package models

// Модель кошелька
type Wallet struct {
	Address string  `json:"address"` // Уникальный адрес кошелька
	Balance float64 `json:"balance"` // Текущий баланс кошелька
}

// Модель транзакции
type Transaction struct {
	ID     int     `json:"id"`     // Уникальный идентификатор транзакции
	From   string  `json:"from"`   // Адрес кошелька отправителя
	To     string  `json:"to"`     // Адрес кошелька получателя
	Amount float64 `json:"amount"` // Сумма перевода
	Time   string  `json:"time"`   // Время выполнения транзакции
}
