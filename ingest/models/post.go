package models

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

var insertPostQuery = `
	INSERT INTO posts(
		  created_at
		, updated_at
		, published_at
		, id
		, collection
		, record_key
		, text
		, actor_id
		, language
		, country
		, locale
		, tags
	) VALUES (
		$1
		, $2
		, $3
		, $4
		, $5
		, $6
		, $7
		, $8
		, $9
		, $10
		, $11
		, $12
	) ON CONFLICT (id) DO NOTHING;
`

type Post struct {
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	PublishedAt *time.Time

	// ID represents the CID for the post
	ID         string
	Collection string
	RecordKey  string
	Text       string
	Actor      *Actor
	Language   string
	Country    string
	Locale     string
	Tags       []string
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
		p.PublishedAt,
		p.ID,
		p.Collection,
		p.RecordKey,
		p.Text,
		p.Actor.ID,
		p.Language,
		p.Country,
		p.Locale,
		pq.Array(p.Tags),
	)

	if err != nil {
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return true, nil
}
