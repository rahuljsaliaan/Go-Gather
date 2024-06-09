package modals

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserId      uint // positive integer
}

var events = []Event{}

func (e Event) Save() {
	// TODO: Save to the data base

	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
