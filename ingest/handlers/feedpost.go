package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/heyts/skylinks/models"
	"github.com/heyts/skylinks/utils"
)

var PostHasNoLinkErr = fmt.Errorf("Post contains no link")

func (h *RecordHandler) FeedPostHandler(op OpMeta, record *bsky.FeedPost) {
	publishedAt, err := time.Parse(time.RFC3339Nano, record.CreatedAt)
	if err != nil {
		h.logger.Error("published date parsing", "value", record.CreatedAt, "err", err)
		return
	}

	createdAt := time.Now().UTC()

	mentions := []*models.Actor{}
	tags := []string{}
	links := []*models.Link{}

	if record.Embed != nil && record.Embed.EmbedRecord != nil && record.Embed.EmbedRecord.Record != nil {
		/* 	If the record has an EmbedRecord, it means that it is a Quote Post, we should just
		   	create a stats record to increment the reposts value
		*/

		q := models.QuotePost{
			CreatedAt: &createdAt,
			PostID:    record.Embed.EmbedRecord.Record.Cid,
		}

		_, err := q.Insert(h.db)
		if err != nil {
			h.logger.Error("quote record", "cid", record.Embed.EmbedRecord.Record.Cid, "err", err)
			return
		}
		h.logger.Info("QUO", "quote", q)
		return
	}

	if len(record.Facets) > 0 {
		for _, f := range record.Facets {
			for _, ft := range f.Features {
				if ft.RichtextFacet_Mention != nil {
					mention, err := h.ParseActorFromDID(ft.RichtextFacet_Mention.Did)
					if err != nil {
						h.logger.Error("actor resolution", "did", ft.RichtextFacet_Mention.Did, "err", err)
						continue
					}
					mentions = append(mentions, mention)
				}

				if ft.RichtextFacet_Tag != nil {
					tag := ft.RichtextFacet_Tag.Tag
					tags = append(tags, tag)
				}

				if ft.RichtextFacet_Link != nil {
					url, meta, err := h.domainResolver.Resolve(ft.RichtextFacet_Link.Uri)
					if err != nil {
						h.logger.Error("metadata resolution", "err", err)
						return
					}

					rawImgopts := map[string]string{}
					rawOpts := map[string]string{}

					if meta != nil {
						for k, v := range *meta {
							if k == "title" ||
								k == "og:title" ||
								k == "og:image" ||
								k == "og:description" ||
								k == "og:site_name" {
								continue
							}

							if strings.HasPrefix(k, "og:image:") {
								rawImgopts[k] = v
							} else {
								rawOpts[k] = v
							}
						}
					}

					imgopts, err := json.Marshal(rawImgopts)
					if err != nil {
						h.logger.Error("JSON Encoding", "err", err)
					}

					opts, err := json.Marshal(rawOpts)
					if err != nil {
						h.logger.Error("JSON Encoding", "err", err)
					}

					link := &models.Link{
						CreatedAt:      &createdAt,
						UpdatedAt:      &createdAt,
						ID:             utils.MD5Encode(url),
						OriginalUrl:    ft.RichtextFacet_Link.Uri,
						Url:            url,
						Title:          meta.GetOrDefault("title", ""),
						OGTitle:        meta.GetOrDefault("og:title", ""),
						OGDescription:  meta.GetOrDefault("og:description", ""),
						OGSiteName:     meta.GetOrDefault("og:site_name", ""),
						OGImage:        meta.GetOrDefault("og:image", ""),
						OGImageOptions: imgopts,
						OGOptional:     opts,
					}
					links = append(links, link)
				}
			}
		}
	}

	// If the record doesn't contain any link
	// we can skip it to preserve resources
	if len(links) == 0 {
		return
	}

	actor, err := h.ParseActorFromDID(op.Repo)
	if err != nil {
		h.logger.Error("parse actor from DID", "err", err)
		return
	}

	var (
		language = ""
		country  = ""
		locale   = ""
	)

	if len(record.Langs) != 0 {
		language = record.Langs[0]
		segments := strings.Split(record.Langs[0], "-")
		country = strings.ToLower(segments[0])
		if len(segments) > 1 {
			locale = segments[1]
		}
	}

	post := models.Post{
		CreatedAt:   &createdAt,
		UpdatedAt:   &createdAt,
		PublishedAt: &publishedAt,

		ID:         op.Cid.String(),
		Collection: op.Collection,
		RecordKey:  op.RecordKey,
		Text:       record.Text,
		Language:   language,
		Country:    country,
		Locale:     locale,
		Tags:       tags,
		Actor:      actor,
	}

	if len(links) > 0 {
		h.logger.Info("POS", "post", post)
		_, err := post.Insert(h.db)
		if err != nil {
			h.logger.Error("insert post", "err", err)
		}

		h.logger.Info("ACT", "actor", actor)
		_, err = actor.Insert(h.db)
		if err != nil {
			h.logger.Error("insert actor", "err", err)
		}

		for _, mention := range mentions {
			h.logger.Debug("Start MEN", "mention", mention)
			_, err = mention.Insert(h.db)
			if err != nil {
				h.logger.Error("insert mention", "err", err)
			}

			_, err = mention.InsertFromPost(h.db, post.ID, mention.ID)
			if err != nil {
				h.logger.Error("insert mention join", "err", err)
			}

			h.logger.Info("MEN", "mention", mention)
		}

		for _, link := range links {
			h.logger.Debug("Start LNK", "link", link)
			_, err = link.Insert(h.db)
			if err != nil {
				h.logger.Error("insert link", "err", err)
			}

			_, err = link.InsertFromPost(h.db, post.ID)
			if err != nil {
				h.logger.Error("insert link join", "err", err)
			}

			h.logger.Info("LNK", "link", link)
		}
	}
}

func (h *RecordHandler) ParseActorFromDID(did string) (*models.Actor, error) {
	ctx := context.Background()
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
		FollowsCount:   *profile.FollowsCount,
		PostsCount:     *profile.PostsCount,
	}

	return actor, nil
}
