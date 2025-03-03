package handlers

import (
	"time"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/heyts/skylinks/models"
)

func (h *RecordHandler) FeedRepostHandler(op OpMeta, record *bsky.FeedRepost) {
	now := time.Now()

	r := &models.Repost{
		CreatedAt: &now,
		UpdatedAt: &now,
		ID:        record.Subject.Cid,
		ActorID:   op.Repo,
	}

	_, err := r.Insert(h.db)
	if err != nil {
		h.logger.Error("repost", "err", err)
		return
	}
	h.logger.Info("RPO", "repost", r)
}
