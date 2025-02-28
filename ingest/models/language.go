package models

import (
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

var insertLanguageQuery = `
	INSERT INTO languages(
		  created_at
		, updated_at
		, id
		, country
		, language
	) VALUES (
		$1
		, $2
		, $3
		, $4
		, $5
	) ON CONFLICT (id) DO NOTHING;
`

var insertLanguageFromPostQuery = `
INSERT INTO posts_languages(
		post_id
		, language_id
	) VALUES (
		$1,$2 
	) ON CONFLICT (post_id, language_id) DO NOTHING;
`

type Language struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time

	// format: <country>[-<language>]
	ID string

	// ISO 3166 Country Code
	Country string

	// Optional ISO 639 Language Code
	Language string
}

func NewLanguage(raw string) *Language {
	if raw == "" {
		return nil
	}

	now := time.Now()
	var lang = &Language{
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	segments := strings.Split(raw, "-")

	if len(segments) == 1 || len(segments) > 2 {
		lang.Country = strings.ToLower(segments[0])
		lang.ID = lang.Country
	}

	if len(segments) == 2 {
		lang.Country = strings.ToLower(segments[0])
		lang.ID = fmt.Sprintf("%s-%s", strings.ToLower(segments[0]), segments[1])
		lang.Language = segments[1]
	}
	return lang
}

func (l Language) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("language", l.ID),
	)
}

func (l *Language) Insert(db *sqlx.DB) (bool, error) {
	tx := db.MustBegin()
	_, err := tx.Exec(insertLanguageQuery,
		l.CreatedAt,
		l.UpdatedAt,
		l.ID,
		l.Country,
		l.Language,
	)

	if err != nil {
		fmt.Printf("err %v", err)
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return true, nil
}

func (l *Language) InsertFromPost(db *sqlx.DB, post_id string) (bool, error) {
	tx := db.MustBegin()
	_, err := tx.Exec(insertLanguageFromPostQuery,
		post_id,
		l.ID,
	)

	if err != nil {
		fmt.Printf("err %v", err)
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return true, nil
}
