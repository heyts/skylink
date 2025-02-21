package ingest

import (
	"github.com/bluesky-social/indigo/api/bsky"
	typegen "github.com/whyrusleeping/cbor-gen"
)

type RecordHandler interface {
	HandleRecord(repo string, record *typegen.CBORMarshaler) error
}

type FeedPost_RecordHandler struct{}

func (r *FeedPost_RecordHandler) HandleRecord(repo string, record *bsky.FeedPost) error {
	return nil
}
