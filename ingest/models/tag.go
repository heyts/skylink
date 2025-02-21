package models

import "time"

type Tag struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time

	// MD5-encoded Tag
	ID string

	// Human-readable Tag
	Label string
}
