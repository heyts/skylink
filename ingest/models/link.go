package models

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
)

var insertLinkQuery = `
	INSERT INTO links(
		  created_at
		, updated_at
		, id
		, url
		, original_url
		, count
		, title
		, og_title
		, og_description
		, og_site_name
		, og_image
		, og_image_options
		, og_optional
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
		, $13

	) ON CONFLICT (id) DO UPDATE SET
	 	updated_at = NOW(),
		count = links.count + 1
`

var insertLinkFromPostQuery = `
INSERT INTO links_posts(
		post_id
		, link_id
	) VALUES (
		$1,$2 
	) ON CONFLICT (post_id, link_id) DO NOTHING;
`

type Link struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time

	ID             string
	OriginalUrl    string
	Url            string
	Count          int
	Title          string
	OGTitle        string
	OGDescription  string
	OGSiteName     string
	OGImage        string
	OGImageOptions []byte
	OGOptional     []byte
}

func (l Link) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("md5", l.ID),
		slog.String("url", l.Url),
	)
}

func (l *Link) Insert(db *sqlx.DB) (bool, error) {
	tx := db.MustBegin()
	_, err := tx.Exec(insertLinkQuery,
		l.CreatedAt,
		l.UpdatedAt,
		l.ID,
		l.Url,
		l.OriginalUrl,
		1,
		l.Title,
		l.OGTitle,
		l.OGDescription,
		l.OGSiteName,
		l.OGImage,
		l.OGImageOptions,
		l.OGOptional,
	)

	if err != nil {
		tx.Rollback()
		fmt.Printf("%v", err)
		return false, err
	}
	tx.Commit()
	return true, nil
}

func (l *Link) InsertFromPost(db *sqlx.DB, post_id string) (bool, error) {
	tx := db.MustBegin()
	_, err := tx.Exec(insertLinkFromPostQuery,
		post_id,
		l.ID,
	)

	if err != nil {
		tx.Rollback()
		fmt.Printf("%v", err)
		return false, err
	}
	tx.Commit()
	return true, nil
}
