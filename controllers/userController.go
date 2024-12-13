package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"user_management/db"
	"user_management/models"
)

// Helper function - error response
func respondWithError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// Helper function - JSON data response
func respondWithJSON(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

// CreateUser
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Insert user -> database
	query := `
		INSERT INTO users (first_name, last_name, email, phone_number, account_type, balance)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`
	var user models.User
	if err := db.DB.QueryRow(
		query,
		req.FirstName,
		req.LastName,
		req.Email,
		req.PhoneNumber,
		req.AccountType,
		req.InitialBalance,
	).Scan(&user.ID, &user.CreatedAt); err != nil {
		respondWithError(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Map the request data to the response model
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.PhoneNumber = req.PhoneNumber
	user.AccountType = req.AccountType
	user.Balance = req.InitialBalance

	respondWithJSON(w, user, http.StatusCreated)
}

// GetUser - by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Fetch user from db
	query := `
		SELECT id, first_name, last_name, email, phone_number, account_type, balance, created_at
		FROM users
		WHERE id = $1
	`
	var user models.User
	if err := db.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PhoneNumber,
		&user.AccountType,
		&user.Balance,
		&user.CreatedAt,
	); err == sql.ErrNoRows {
		respondWithError(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		respondWithError(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, user, http.StatusOK)
}

// UpdateUser
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update user in database
	query := `
		UPDATE users 
		SET first_name = $1, last_name = $2, email = $3, 
		    phone_number = $4, account_type = $5
		WHERE id = $6
	`
	if _, err := db.DB.Exec(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PhoneNumber,
		user.AccountType,
		user.ID,
	); err != nil {
		respondWithError(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteUser - by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Delete user from database
	query := `DELETE FROM users WHERE id = $1`
	if _, err := db.DB.Exec(query, id); err != nil {
		respondWithError(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// ListUsers - retrieves all users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT id, first_name, last_name, email, phone_number, 
		       account_type, balance, created_at
		FROM users
		ORDER BY created_at DESC
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		respondWithError(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.PhoneNumber,
			&user.AccountType,
			&user.Balance,
			&user.CreatedAt,
		); err != nil {
			respondWithError(w, "Error scanning user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	respondWithJSON(w, users, http.StatusOK)
}
