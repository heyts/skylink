package models

import "time"

type Actor struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time

	ID          string
	DisplayName string
	Handle      string
}
