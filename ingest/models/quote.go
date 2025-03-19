package models

import (
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
)

type QuotePost struct {
	CreatedAt *time.Time
	PostID    string
}

func (q QuotePost) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", q.PostID),
	)
}

func (q *QuotePost) Insert(db *sqlx.DB) (bool, error) {
	ps := TimeRangeStat{
		YMDH:        q.CreatedAt,
		PostID:      q.PostID,
		QuotesCount: 1,
	}

	tx := db.MustBegin()
	_, err := ps.InsertMultiple(db, []string{"hour", "day", "week", "month"})
	if err != nil {
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return true, nil
}
