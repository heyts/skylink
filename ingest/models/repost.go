package models

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
)

var insertRepostQuery = `
	INSERT INTO reposts(
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
	-- ON CONFLICT (id, actor_id) DO NOTHING;
`

type Repost struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time

	// ID represents the CID for the post
	ID      string
	ActorID string
}

func (p Repost) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", p.ID),
	)
}

func (p *Repost) Insert(db *sqlx.DB) (bool, error) {
	tx := db.MustBegin()
	_, err := tx.Exec(insertRepostQuery,
		p.CreatedAt,
		p.UpdatedAt,
		p.ID,
		p.ActorID,
	)

	if err != nil {
		fmt.Printf("err: %v", err)
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return true, nil
}
