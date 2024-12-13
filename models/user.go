package models

import (
	"time"
)

type User struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	AccountType string    `json:"account_type"`
	Balance     float64   `json:"balance"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateUserRequest struct {
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	Email          string  `json:"email"`
	PhoneNumber    string  `json:"phone_number"`
	AccountType    string  `json:"account_type"`
	InitialBalance float64 `json:"initial_balance"`
}
