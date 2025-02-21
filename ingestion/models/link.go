package models

import "time"

type Link struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time

	ID  string
	Url string
}
