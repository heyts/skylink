package models

import (
	"bytes"
	"fmt"
	"html/template"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
)

// 'hourly_stats': 'hour'
// 'daily_stats': 'day'
// 'weekly_stats': 'week'
// 'monthly_stats': 'month'

var validTimeRanges = map[string]string{
	"hour":  "hourly_stats",
	"day":   "daily_stats",
	"week":  "weekly_stats",
	"month": "monthly_stats",
}

// This query inserts a new record if no record exist in stats and if
// a post record matching the post_id exist in posts, otherwise if no match is found
// the record is ignored. If a match is found but a record already exist in stats
// the record counts are updated
var insertTimeRangeStatQuery = `
	INSERT INTO {{ .tablename }} (
	ymdh, 
	post_id, 
	likes_count, 
	reposts_count,
	quotes_count
	) SELECT 
	 	v.ymdh, 
		v.post_id, 
		v.likes_count, 
		v.reposts_count,
		v.quotes_count 
	FROM (
		VALUES (
			date_trunc('{{ .period }}', $1::timestamp), 
			$2, 
			$3::integer, 
			$4::integer,
			$5::integer
		) 
	) v (ymdh, post_id, likes_count, reposts_count, quotes_count)
	JOIN posts p on p.id = v.post_id 
	ON CONFLICT (ymdh, post_id) DO UPDATE SET
		likes_count = {{ .tablename }}.likes_count + excluded.likes_count,
		quotes_count = {{ .tablename }}.quotes_count + excluded.quotes_count,
        reposts_count = {{ .tablename }}.reposts_count + excluded.reposts_count;
`

var tpl = template.Must(template.New("TimeRangeStatQuery").Parse(insertTimeRangeStatQuery))

type TimeRangeStat struct {
	YMDH *time.Time

	// ID represents the CID for the post
	PostID       string
	LikesCount   int64
	RepostsCount int64
	QuotesCount  int64
}

func (tr TimeRangeStat) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", tr.PostID),
	)
}

func (tr *TimeRangeStat) Insert(db *sqlx.DB, period string) (bool, error) {
	query, err := tr.TimeRangeQuery(period)
	if err != nil {
		return false, err
	}

	tx := db.MustBegin()
	_, err = tx.Exec(query,
		tr.YMDH,
		tr.PostID,
		tr.LikesCount,
		tr.RepostsCount,
		tr.QuotesCount,
	)

	if err != nil {
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return true, nil
}

func (tr *TimeRangeStat) InsertMultiple(db *sqlx.DB, periods []string) (bool, error) {
	for _, p := range periods {
		query, err := tr.TimeRangeQuery(p)
		if err != nil {
			return false, err
		}

		tx := db.MustBegin()
		_, err = tx.Exec(query,
			tr.YMDH,
			tr.PostID,
			tr.LikesCount,
			tr.RepostsCount,
			tr.QuotesCount,
		)

		if err != nil {
			tx.Rollback()
			return false, err
		}
		tx.Commit()
	}
	return true, nil
}

func (p *TimeRangeStat) TimeRangeQuery(period string) (string, error) {
	table, ok := validTimeRanges[period]
	if !ok {
		return "", fmt.Errorf("invalid period: %q", period)
	}

	m := make(map[string]string)
	m["period"] = period
	m["tablename"] = table

	parsed := bytes.NewBufferString("")

	err := tpl.Execute(parsed, m)
	if err != nil {
		return "", err
	}

	return parsed.String(), nil

}
