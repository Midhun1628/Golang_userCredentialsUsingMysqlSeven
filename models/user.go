package models

import (
	"database/sql"
	"log"
)

// User struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Authenticate checks if user exists in database
func Authenticate(db *sql.DB, username, password string) (bool, error) {
	var user User
	err := db.QueryRow("SELECT username, password FROM credentials WHERE username=? AND password=?", username, password).Scan(&user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // User not found
		}
		return false, err // Some other error
	}
	return true, nil
}

// GetAllUsers retrieves all users from the database
func GetAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT username, password FROM credentials")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Username, &user.Password); err != nil {
			log.Println("Error scanning user:", err)
			continue
		}
		users = append(users, user)
	}

	return users, nil
}
