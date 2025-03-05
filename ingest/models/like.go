package models

import (
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
)

type Like struct {
	CreatedAt *time.Time
	PostID    string
}

func (l Like) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", l.PostID),
	)
}

func (l *Like) Insert(db *sqlx.DB) (bool, error) {
	tx := db.MustBegin()

	ps := TimeRangeStat{
		YMDH:         l.CreatedAt,
		PostID:       l.PostID,
		RepostsCount: 0,
		LikesCount:   1,
	}

	_, err := ps.InsertMultiple(db, []string{"hour", "day", "week", "month"})
	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	return true, nil
}
