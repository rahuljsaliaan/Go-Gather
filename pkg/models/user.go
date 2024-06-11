package models

import "rahuljsaliaan.com/go-gather/internal/db"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := `
		INSERT INTO users(email, password)
		VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Email, user.Password)

	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()

	if err != nil {
		return err
	}

	user.ID = userID

	return nil
}
