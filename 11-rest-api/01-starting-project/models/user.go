package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "insert into users(email, password) values(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashed, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hashed)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	u.ID = id

	return err
}

func (u *User) ValidateCredentials() error {
	query := "select id, password from users where email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var password string
	err := row.Scan(&u.ID, &password)
	if err != nil {
		return err
	}

	isValid := utils.CheckPassword(u.Password, password)

	if !isValid {
		return errors.New("invalid password")
	}

	return nil
}
