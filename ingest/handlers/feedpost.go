package handlers

import (
	"context"
	"time"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/heyts/skylinks/models"
)

func (h *RecordHandler) FeedPostHandler(op OpMeta, record *bsky.FeedPost) error {
	// actor, err := h.resolveUserIdentity(op.Repo)
	actor, err := bsky.ActorGetProfile(context.Background(), h.xrpcClient, op.Repo)
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

	model := &models.HyperLink{
		CreatedAt:  &createdAt,
		CID:        op.Cid,
		Collection: op.Collection,
		RecordKey:  op.RecordKey,
		Actor:      actor,
		Languages:  record.Langs,
		Text:       record.Text,
		Mentions:   mentions,
		Tags:       tags,
		URI:        uris,
	}

	// sqlite3.

	if len(model.URI) > 0 {
		h.logger.Info("record", "val", model)
	}
	return nil
}
