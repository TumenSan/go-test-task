package database

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDatabase() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS wallets (
			address TEXT PRIMARY KEY,
			balance REAL NOT NULL
	)`)

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS transactions (
			id SERIAL PRIMARY KEY,
			from_wallet TEXT,
			to_wallet TEXT,
			amount REAL,
			time TIMESTAMP
	)`)

	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	seedWallets()
}

func seedWallets() {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM wallets").Scan(&count)
	if err != nil {
		log.Fatalf("Failed to check wallets count: %v", err)
	}

	if count == 0 {
		for i := 0; i < 10; i++ {
			address := randomAddress()
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

func randomAddress() string {
	return fmt.Sprintf("%x", rand.Int63())
}
