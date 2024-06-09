package models

import "time"

type Event struct {
	ID          uint
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      uint      // positive integer
}

var events = []Event{}

func (e Event) Save() {
	// TODO: Save to the data base

	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
