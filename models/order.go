package models

import "time"

type Order struct {
	Type     string    `json:"type" db:"type"`
	Cost     int       `json:"cost" db:"cost"`
	Time     time.Time `json:"time" db:"time"`
	Duration int       `json:"duration" db:"duration"`
}
