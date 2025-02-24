package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/heyts/skylinks/models"
	"github.com/heyts/skylinks/utils"
)

func (h *RecordHandler) FeedPostHandler(op OpMeta, record *bsky.FeedPost) error {
	// actor, err := h.resolveUserIdentity(op.Repo)
	profile, err := bsky.ActorGetProfile(context.Background(), h.xrpcClient, op.Repo)
	if err != nil {
		return err
	}

	mentions := []string{}
	tags := []string{}
	uris := []string{}

	if len(record.Facets) > 0 {
		for _, f := range record.Facets {
			for _, ft := range f.Features {
				if ft.RichtextFacet_Mention != nil {
					mentions = append(mentions, ft.RichtextFacet_Mention.Did)
				}

				if ft.RichtextFacet_Tag != nil {
					tags = append(tags, ft.RichtextFacet_Tag.Tag)
				}
				if ft.RichtextFacet_Link != nil {
					url, err := h.domainResolver.Resolve(ft.RichtextFacet_Link.Uri)
					if err != nil {
						return err
					}
					uris = append(uris, url)
				}
			}
		}
	}
	createdAt, err := time.Parse(time.RFC3339Nano, record.CreatedAt)
	if err != nil {
		return err
	}

	post := models.Post{
		CreatedAt:  &createdAt,
		UpdatedAt:  &createdAt,
		ID:         op.Cid.String(),
		Collection: op.Collection,
		RecordKey:  op.RecordKey,
		Text:       record.Text,
		ActorID:    profile.Did,
	}

	actor := models.Actor{
		CreatedAt:   &createdAt,
		UpdatedAt:   &createdAt,
		ID:          profile.Did,
		DisplayName: *profile.DisplayName,
		Handle:      profile.Handle,
	}

	links := []models.Link{}

	for _, uri := range uris {
		link := models.Link{
			CreatedAt: &createdAt,
			UpdatedAt: &createdAt,
			ID:        utils.MD5Encode(uri),
			Url:       uri,
		}
		links = append(links, link)
	}

	// 1.actor,
	// 2.post
	// 3.language,
	// 4.mentions,
	// 5.tags,
	// 6.uris
	/*
		model := &models.HyperLink{
			CreatedAt:  &createdAt,
			CID:        op.Cid,
			Collection: op.Collection,
			RecordKey:  op.RecordKey,
			Actor:      profile,
			Languages:  record.Langs,
			Text:       record.Text,
			Mentions:   mentions,
			Tags:       tags,
			URI:        uris,
		}
	*/

	// sqlite3.

	if len(links) > 0 {
		fmt.Printf("Post: %+v\n", post)
		fmt.Printf("Actor: %+v\n", actor)
		for _, link := range links {
			fmt.Printf("Link: %+v\n\n", link)
		}
	}
	return nil
}
