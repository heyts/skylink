package handlers

import (
	"log/slog"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/xrpc"
	"github.com/heyts/skylinks/utils"
	"github.com/jmoiron/sqlx"
)

type RecordHandler struct {
	logger         *slog.Logger
	db             *sqlx.DB
	domainResolver *utils.DomainResolver
	xrpcClient     *xrpc.Client
	workChan       <-chan OpMeta
	quitChan       <-chan struct{}
	errs           map[string]string
}

func NewRecordHandler(workChan <-chan OpMeta, quitChan <-chan struct{}, logger *slog.Logger, db *sqlx.DB, domainResolver *utils.DomainResolver) {
	r := &RecordHandler{
		logger:         logger,
		db:             db,
		domainResolver: domainResolver,
		xrpcClient: &xrpc.Client{
			Host: "https://public.api.bsky.app",
		},
		workChan: workChan,
		quitChan: quitChan,
		errs:     map[string]string{},
	}

	for {
		select {
		case op := <-workChan:
			switch record := op.Record.(type) {
			case *bsky.FeedPost:
				r.FeedPostHandler(op, record)

			case *bsky.FeedRepost:
				r.FeedRepostHandler(op, record)

			case *bsky.FeedLike:
				r.FeedLikeHandler(op, record)
			}

		case <-quitChan:
			return
		}
	}
}
