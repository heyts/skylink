package models

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
)

var insertActorQuery = `
	INSERT INTO actors(
		  created_at
		, updated_at
		, id
		, display_name
		, handle
		, avatar
		, banner
		, followers_count
		, follows_count
		, posts_count
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
	) ON CONFLICT (id) DO UPDATE SET
	 	updated_at = NOW(),
		followers_count = excluded.followers_count,
		follows_count = excluded.follows_count,
		posts_count = excluded.posts_count;
`

var insertActorFromMentionQuery = `
	INSERT INTO mentions_posts(
		post_id
		, actor_id
	) VALUES (
		$1,$2 
	) ON CONFLICT (post_id, actor_id) DO NOTHING;
`

type Actor struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time

	ID             string
	DisplayName    string
	Handle         string
	Avatar         string
	Banner         string
	FollowersCount int64
	FollowsCount   int64
	PostsCount     int64
}

func (a Actor) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", a.ID),
		slog.String("handle", a.Handle),
	)
}

func (a Actor) URL() string {
	return fmt.Sprintf("https://bsky.app/profile/%s", a.Handle)
}

func (a *Actor) Insert(db *sqlx.DB) (bool, error) {
	tx := db.MustBegin()
	_, err := tx.Exec(insertActorQuery,
		a.CreatedAt,
		a.UpdatedAt,
		a.ID,
		a.DisplayName,
		a.Handle,
		a.Avatar,
		a.Banner,
		a.FollowersCount,
		a.FollowsCount,
		a.PostsCount,
	)

	if err != nil {
		tx.Rollback()
		fmt.Printf("%v", err)
		return false, err
	}
	tx.Commit()
	return true, nil
}

func (a *Actor) InsertFromPost(db *sqlx.DB, post_id, actor_id string) (bool, error) {
	tx := db.MustBegin()
	_, err := tx.Exec(insertActorFromMentionQuery,
		post_id,
		actor_id,
	)

	if err != nil {
		tx.Rollback()
		fmt.Printf("%v", err)
		return false, err
	}
	tx.Commit()
	return true, nil
}

func (a *Actor) Exists(db *sqlx.DB, id string) bool {
	return true
}
