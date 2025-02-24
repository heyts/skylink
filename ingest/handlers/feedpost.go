package handlers

import (
	"context"
	"time"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/heyts/skylinks/models"
	"github.com/heyts/skylinks/utils"
)

func (h *RecordHandler) FeedPostHandler(op OpMeta, record *bsky.FeedPost) error {
	actor, err := h.ParseActorFromDID(op.Repo)
	if err != nil {
		return err
	}
	createdAt, err := time.Parse(time.RFC3339Nano, record.CreatedAt)

	post := models.Post{
		CreatedAt:  &createdAt,
		UpdatedAt:  &createdAt,
		ID:         op.Cid.String(),
		Collection: op.Collection,
		RecordKey:  op.RecordKey,
		Text:       record.Text,
		Actor:      actor,
	}

	languages := []*models.Language{}
	if len(record.Langs) > 0 {
		for _, lang := range record.Langs {
			lang := models.NewLanguage(lang)
			if lang != nil {
				languages = append(languages, lang)
			}
		}
	}

	mentions := []*models.Actor{}
	tags := []*models.Tag{}
	links := []*models.Link{}

	if len(record.Facets) > 0 {
		for _, f := range record.Facets {
			for _, ft := range f.Features {
				if ft.RichtextFacet_Mention != nil {
					mention, err := h.ParseActorFromDID(op.Repo)
					if err != nil {
						h.logger.Error("actor resolution", "did", ft.RichtextFacet_Mention.Did, "err", err)
						continue
					}
					mentions = append(mentions, mention)
				}

				if ft.RichtextFacet_Tag != nil {
					tag := &models.Tag{
						CreatedAt: &createdAt,
						UpdatedAt: &createdAt,
						ID:        utils.MD5Encode(ft.RichtextFacet_Tag.Tag),
						Label:     ft.RichtextFacet_Tag.Tag,
					}
					tags = append(tags, tag)
				}

				if ft.RichtextFacet_Link != nil {
					url, err := h.domainResolver.Resolve(ft.RichtextFacet_Link.Uri)
					if err != nil {
						return err
					}
					link := &models.Link{
						CreatedAt: &createdAt,
						UpdatedAt: &createdAt,
						ID:        utils.MD5Encode(url),
						Url:       url,
					}
					links = append(links, link)
				}
			}
		}
	}

	if len(links) > 0 {
		h.logger.Info("POS", "post", post)
		go post.Insert(h.db)
		if err != nil {
			h.logger.Error("insert post", "err", err)
		}

		h.logger.Info("ACT", "actor", actor)
		_, err = actor.Insert(h.db)
		if err != nil {
			h.logger.Error("insert actor", "err", err)
		}

		if len(languages) > 0 {
			for _, lang := range languages {
				h.logger.Info("LNG", "language", lang)
				go lang.Insert(h.db)
				if err != nil {
					h.logger.Error("insert lang", "err", err)
				}
				go lang.InsertFromPost(h.db, post.ID)
				if err != nil {
					h.logger.Error("insert language join table", "err", err)
				}
			}
		}

		for _, tag := range tags {
			h.logger.Info("TAG", "tag", tag)
			go tag.Insert(h.db)
			if err != nil {
				h.logger.Error("insert tag", "err", err)
			}
			go tag.InsertFromPost(h.db, post.ID)
			if err != nil {
				h.logger.Error("insert tag join table", "err", err)
			}
		}

		for _, mention := range mentions {
			go mention.Insert(h.db)
			if err != nil {
				h.logger.Error("insert mention", "err", err)
			}

			go mention.InsertFromPost(h.db, post.ID)
			if err != nil {
				h.logger.Error("insert mention join table", "err", err)
			}

			h.logger.Info("MEN", "mention", mention)
		}

		for _, link := range links {
			go link.Insert(h.db)
			if err != nil {
				h.logger.Error("insert link", "err", err)
			}

			go link.InsertFromPost(h.db, post.ID)
			if err != nil {
				h.logger.Error("insert link join table", "err", err)
			}
			h.logger.Info("LNK", "link", link)
		}
	}
	return nil
}

func (h *RecordHandler) ParseActorFromDID(did string) (*models.Actor, error) {
	ctx, _ := context.WithCancel(context.Background())
	profile, err := bsky.ActorGetProfile(ctx, h.xrpcClient, did)
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339Nano, *profile.CreatedAt)
	if err != nil {
		return nil, err
	}

	if profile.DisplayName == nil {
		e := new(string)
		profile.DisplayName = e
	}

	if profile.Avatar == nil {
		e := new(string)
		profile.Avatar = e
	}

	if profile.Banner == nil {
		e := new(string)
		profile.Banner = e
	}

	if profile.FollowersCount == nil {
		e := new(int64)
		profile.FollowersCount = e
	}

	if profile.FollowsCount == nil {
		e := new(int64)
		profile.FollowsCount = e
	}

	if profile.PostsCount == nil {
		e := new(int64)
		profile.PostsCount = e
	}

	actor := &models.Actor{
		CreatedAt:      &createdAt,
		UpdatedAt:      &createdAt,
		ID:             profile.Did,
		DisplayName:    *profile.DisplayName,
		Handle:         profile.Handle,
		Avatar:         *profile.Avatar,
		Banner:         *profile.Banner,
		FollowersCount: *profile.FollowersCount,
		FollowsCount:   *profile.FollowersCount,
		PostsCount:     *profile.PostsCount,
	}

	return actor, nil
}
