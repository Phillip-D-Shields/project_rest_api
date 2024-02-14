package models

import (
	"errors"

	"example.com/rest-project/db"
	"example.com/rest-project/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`

	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.Db.QueryRow(query, &u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("invalid creds")
	}

	passwordIsValid := utils.VerifyPassword(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid creds")
	}

	return nil
}
