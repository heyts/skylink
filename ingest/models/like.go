package models

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
)

var insertLikeQuery = `
	INSERT INTO likes(
		  created_at
		, updated_at
		, id
		, actor_id
	) VALUES (
		$1
		, $2
		, $3
		, $4
	); 
`

type Like struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time

	// ID represents the CID for the post
	ID      string
	ActorID string
}

func (l Like) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", l.ID),
	)
}

func (l *Like) Insert(db *sqlx.DB) (bool, error) {
	tx := db.MustBegin()
	_, err := tx.Exec(insertLikeQuery,
		l.CreatedAt,
		l.UpdatedAt,
		l.ID,
		l.ActorID,
	)

	if err != nil {
		fmt.Printf("err: %v", err)
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return true, nil
}
