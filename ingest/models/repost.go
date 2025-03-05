package models

import (
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
)

type Repost struct {
	CreatedAt *time.Time
	PostID    string
}

func (p Repost) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", p.PostID),
	)
}

func (p *Repost) Insert(db *sqlx.DB) (bool, error) {
	tx := db.MustBegin()

	ps := TimeRangeStat{
		YMDH:         p.CreatedAt,
		PostID:       p.PostID,
		RepostsCount: 1,
		LikesCount:   0,
	}

	_, err := ps.InsertMultiple(db, []string{"hour", "day", "week", "month"})
	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	return true, nil
}
