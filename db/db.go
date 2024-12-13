// db/connection.go
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDatabase initializes the db connection
func InitDatabase() error {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	// Database connection params
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open db connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database connection: %w", err)
	}

	// Testing the connection
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error testing database connection: %w", err)
	}

	DB = db
	log.Println("Database connection established")

	// Run migrations
	if err := createTables(); err != nil {
		return fmt.Errorf("error creating tables: %w", err)
	}

	return nil
}

// createTables creates the necessary tables in the db
func createTables() error {
	// Create users table if it does not exist
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			first_name VARCHAR(100) NOT NULL,
			last_name VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			phone_number VARCHAR(20),
			account_type VARCHAR(50) NOT NULL,
			balance DECIMAL(15,2) DEFAULT 0.00,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("error creating users table: %w", err)
	}

	return nil
}

// CloseDatabase closes the db connection
func CloseDatabase() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
