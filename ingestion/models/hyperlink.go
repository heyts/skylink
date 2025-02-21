package models

import (
	"fmt"
	"time"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/ipfs/go-cid"
)

type HyperLink struct {
	CreatedAt  *time.Time
	CID        cid.Cid
	Collection string
	RecordKey  string
	Text       string
	Actor      *bsky.ActorDefs_ProfileView
	Languages  []string
	Mentions   []string
	Tags       []string
	URI        []string
}

func (h *HyperLink) String() string {
	return fmt.Sprintf("u=%s coll=%s key=%s lg=%v mn=%v ht=%v u=%+v pu=%s", h.Actor.Handle, h.Collection, h.RecordKey, h.Languages, h.Mentions, h.Tags, h.URI, h.URL())
}

func (h *HyperLink) URL() string {
	return fmt.Sprintf("https://bsky.app/profile/%s/post/%s", h.Actor.Handle, h.RecordKey)
}
