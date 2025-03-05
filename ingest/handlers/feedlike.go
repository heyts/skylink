package handlers

import (
	"time"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/heyts/skylinks/models"
)

func (h *RecordHandler) FeedLikeHandler(op OpMeta, record *bsky.FeedLike) {
	now := time.Now().UTC()

	r := &models.Like{
		CreatedAt: &now,
		PostID:    record.Subject.Cid,
	}

	_, err := r.Insert(h.db)
	if err != nil {
		h.logger.Error("like", "err", err)
		return
	}
	h.logger.Info("LKE", "like", r)
}
