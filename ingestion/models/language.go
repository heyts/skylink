package models

import "time"

type Language struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time

	// format: <country>[-<language>]
	ID string

	// ISO 3166 Country Code
	Country string

	// Optional ISO 639 Language Code
	Language string
}
