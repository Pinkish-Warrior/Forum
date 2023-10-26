package forum

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func NewUser(email, username, password string, db *sql.DB) error {
	// Generate a hashed password from the provided password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Insert the new user into the database
	_, err = db.Exec("INSERT INTO users (email, username, password) VALUES (?, ?, ?)", email, username, string(hashedPassword))
	return err
}
