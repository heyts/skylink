package models

import "time"

type Post struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time

	// ID represents the CID for the post
	ID         string
	Collection string
	RecordKey  string
	Text       string
	ActorID    string
}
