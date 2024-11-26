package models

import (
	"context"
	"go-backend/utils"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

// GetUserByUsername finds a user by username
func GetUserByUsername(username string) (*User, error) {
	var user User
	query := `SELECT id, username, password, created_at FROM users WHERE username=$1`
	err := utils.DB.QueryRow(context.Background(), query, username).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser adds a new user to the database
func CreateUser(username, password string) error {
	query := `INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3)`
	_, err := utils.DB.Exec(context.Background(), query, username, password, time.Now())
	return err
}

// DeletedUser from the database by ID
func DeleteUser(userID int) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := utils.DB.Exec(context.Background(), query, userID)
	return err
}