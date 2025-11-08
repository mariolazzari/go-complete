package models

import "example.com/rest-api/db"

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

	res, err := stmt.Exec(u.Email, u.Email)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	u.ID = id

	return err
}
