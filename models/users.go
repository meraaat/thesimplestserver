package models

import (
	"errors"

	"server.example.com/db"
	"server.example.com/utilities"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPass, err := utilities.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPass)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {

	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPass string
	if err := row.Scan(&u.ID, &retrievedPass); err != nil {
		return err
	}

	isValidPassword := utilities.IsHashPasswordMatch(u.Password, retrievedPass)

	if !isValidPassword {
		return errors.New("credentials invalid")
	}

	return nil
}
