package models

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
)

var insertPostQuery = `
	INSERT INTO posts(
		  created_at
		, updated_at
		, id
		, collection
		, record_key
		, text
		, actor_id
	) VALUES (
		$1
		, $2
		, $3
		, $4
		, $5
		, $6
		, $7
	) ON CONFLICT (id) DO NOTHING;
`

type Post struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time

	// ID represents the CID for the post
	ID         string
	Collection string
	RecordKey  string
	Text       string
	Actor      *Actor
	handle     string
}

func (p Post) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", p.ID),
		slog.String("url", p.URL()),
		slog.String("actor", p.Actor.Handle),
	)
}

func (p Post) URL() string {
	return fmt.Sprintf("https://bsky.app/profile/%s/post/%s", p.Actor.Handle, p.RecordKey)
}

func (p *Post) Insert(db *sqlx.DB) (bool, error) {
	tx := db.MustBegin()
	_, err := tx.Exec(insertPostQuery,
		p.CreatedAt,
		p.UpdatedAt,
		p.ID,
		p.Collection,
		p.RecordKey,
		p.Text,
		p.Actor.ID,
	)

	if err != nil {
		fmt.Printf("err: %v", err)
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return true, nil
}
