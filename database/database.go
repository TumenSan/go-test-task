package database

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"

	_ "github.com/lib/pq" // Драйвер для PostgreSQL
)

// DB — глобальная переменная для хранения подключения к базе данных
var DB *sql.DB

// InitDatabase инициализирует подключение к базе данных и создает таблицы, если они не существуют
// Также выполняет начальное заполнение таблицы wallets тестовыми данными, если она пуста
func InitDatabase() {
	var err error

	// Формирование строки подключения к базе данных из переменных окружения
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	// Подключение к базе данных
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Создание таблицы wallets, если она не существует
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS wallets (
			address TEXT PRIMARY KEY,
			balance REAL NOT NULL
	)`)

	// Создание таблицы transactions, если она не существует
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS transactions (
			id SERIAL PRIMARY KEY,
			from_wallet TEXT,
			to_wallet TEXT,
			amount REAL,
			time TIMESTAMP
	)`)

	// Проверка ошибок при создании таблиц
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	// Заполнение таблицы wallets тестовыми данными, если она пуста
	seedWallets()
}

// seedWallets заполняет таблицу wallets тестовыми данными, если она пуста
// Создает 10 кошельков с начальным балансом 100.0
func seedWallets() {
	var count int

	// Проверка количества записей в таблице wallets
	err := DB.QueryRow("SELECT COUNT(*) FROM wallets").Scan(&count)
	if err != nil {
		log.Fatalf("Failed to check wallets count: %v", err)
	}

	// Если таблица пуста, заполняем её тестовыми данными
	if count == 0 {
		for i := 0; i < 10; i++ {
			address := randomAddress() // Генерация случайного адреса кошелька
			_, err := DB.Exec("INSERT INTO wallets (address, balance) VALUES ($1, $2)", address, 100.0)
			if err != nil {
				log.Fatalf("Failed to seed wallets: %v", err)
			}
		}
		log.Println("Wallets seeded successfully")
	} else {
		log.Println("Wallets already exist, skipping seeding")
	}
}

// randomAddress генерирует случайный адрес кошелька
// Возвращает строку, представляющую случайное шестнадцатеричное число
func randomAddress() string {
	return fmt.Sprintf("%x", rand.Int63())
}
