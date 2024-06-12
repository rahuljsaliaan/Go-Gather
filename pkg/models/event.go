package models

import (
	"errors"
	"fmt"
	"time"

	"rahuljsaliaan.com/go-gather/internal/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64     // positive integer
}

func (event *Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, datetime, user_id)
		VALUES (?, ?, ?, ?, ?)
	 `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	event.ID = id

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Location, &event.UserID)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, err
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events SET name = ?, description = ?, location = ?, datetime = ?
	WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)

	return err
}

func (event Event) Register(userID int64) error {
	query := "INSERT INTO registrations (event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID, userID)

	return err
}

func (event Event) CancelRegistration(userID int64) error {
	fmt.Println(event.ID, userID)
	query := "DELETE FROM registrations WHERE event_id = ? and user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(event.ID, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows were deleted")
	}

	return err
}
