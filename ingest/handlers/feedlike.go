package handlers

import (
	"time"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/heyts/skylinks/models"
)

func (h *RecordHandler) FeedLikeHandler(op OpMeta, record *bsky.FeedLike) {
	now := time.Now()

	r := &models.Like{
		CreatedAt: &now,
		UpdatedAt: &now,
		ID:        record.Subject.Cid,
		ActorID:   op.Repo,
	}

	_, err := r.Insert(h.db)
	if err != nil {
		h.logger.Error("like", "err", err)
		return
	}
	h.logger.Info("LKE", "like", r)
}
