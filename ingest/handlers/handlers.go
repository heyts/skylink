package handlers

import (
	"log/slog"

	"github.com/bluesky-social/indigo/xrpc"
	"github.com/heyts/skylinks/utils"
	"github.com/jmoiron/sqlx"
)

type RecordHandler struct {
	logger         *slog.Logger
	db             *sqlx.DB
	domainResolver *utils.DomainResolver
	xrpcClient     *xrpc.Client
}

func NewRecordHandler(logger *slog.Logger, db *sqlx.DB, domainResolver *utils.DomainResolver) *RecordHandler {
	return &RecordHandler{
		logger:         logger,
		db:             db,
		domainResolver: domainResolver,
		xrpcClient: &xrpc.Client{
			Host: "https://public.api.bsky.app",
		},
	}
}
