package main

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system.
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // The '-' tag prevents the password from being sent in JSON responses
	CreatedAt time.Time `json:"createdAt"`
}

// HashPassword hashes the user's password using bcrypt.
func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword checks if the provided password matches the hashed password.
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}