package models

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
)

var insertTagQuery = `
	INSERT INTO tags(
		  created_at
		, updated_at
		, id
		, label
	) VALUES (
		$1
		, $2
		, $3
		, $4
	) ON CONFLICT (id) DO UPDATE SET
	 	updated_at = NOW();
`

var insertTagFromPostQuery = `
INSERT INTO posts_tags(
		post_id
		, tag_id
	) VALUES (
		$1,$2 
	) ON CONFLICT (post_id, tag_id) DO NOTHING;
`

type Tag struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time

	// MD5-encoded Tag
	ID string

	// Human-readable Tag
	Label string
}

func (t Tag) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("md5", t.ID),
		slog.String("label", t.Label),
	)
}

func (t *Tag) Insert(db *sqlx.DB) (bool, error) {
	tx := db.MustBegin()
	_, err := tx.Exec(insertTagQuery,
		t.CreatedAt,
		t.UpdatedAt,
		t.ID,
		t.Label,
	)

	if err != nil {
		tx.Rollback()
		fmt.Printf("%v", err)
		return false, err
	}
	tx.Commit()
	return true, nil
}

func (t *Tag) InsertFromPost(db *sqlx.DB, post_id string) (bool, error) {
	tx := db.MustBegin()
	_, err := tx.Exec(insertTagFromPostQuery,
		post_id,
		t.ID,
	)

	if err != nil {
		tx.Rollback()
		fmt.Printf("%v", err)
		return false, err
	}
	tx.Commit()
	return true, nil
}
