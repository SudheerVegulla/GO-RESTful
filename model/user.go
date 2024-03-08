package model

import (
	"errors"

	"example.com/event-bboking/db"
	"example.com/event-bboking/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users (email,password) VALUES(?,?)`
	stmt, err := db.DB.Prepare(query)

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
	userId, err := result.LastInsertId()
	u.Id = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials Invalid!")
	}

	passwordIsValid := utils.CompareCredentials(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials Invalid!")
	}

	return nil

}
