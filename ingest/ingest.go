package ingest

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/events"
	"github.com/bluesky-social/indigo/events/schedulers/sequential"
	"github.com/bluesky-social/indigo/repo"
	"github.com/gorilla/websocket"
	"github.com/heyts/skylinks/handlers"
	"github.com/heyts/skylinks/utils"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var resolverPolicy = map[string][]string{
	"youtube.com":     {"v"},
	"ycombinator.com": {"id"},
	"bsky.app":        {"q"},
	"instagram.com":   {"next"},
	"facebook.com":    {"id", "story_fbid", "next"},
	"google.com":      {"q"},
}

type Server struct {
	logger         *slog.Logger
	wsEndpoint     string
	dsn            string
	recordHandlers *handlers.RecordHandler
}

func NewServer(dsn *string) *Server {
	db, err := sqlx.Open("postgres", *dsn)
	if err != nil {
		log.Fatal(err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	rh := handlers.NewRecordHandler(logger, db, utils.NewDomainResolver(resolverPolicy))

	s := &Server{
		logger:         logger,
		wsEndpoint:     "wss://bsky.network/xrpc/com.atproto.sync.subscribeRepos",
		dsn:            *dsn,
		recordHandlers: rh,
	}

	return s
}

func (s *Server) Start() error {
	con, _, err := websocket.DefaultDialer.Dial(s.wsEndpoint, http.Header{})
	if err != nil {
		s.logger.Error(fmt.Sprintf("Could not connect to %q: %v\n", s.wsEndpoint, err))
		os.Exit(1)
	}

	rsc := &events.RepoStreamCallbacks{
		RepoCommit: func(evt *atproto.SyncSubscribeRepos_Commit) error {
			for _, op := range evt.Ops {
				ctx := context.Background()

				raw, err := repo.ReadRepoFromCar(ctx, bytes.NewReader(evt.Blocks))
				if err != nil {
					s.logger.Debug("ReadRepoFromCar", "msg", err)
				}

				cid, rec, err := raw.GetRecord(ctx, op.Path)
				if err != nil {
					s.logger.Debug("GetRecord", "msg", err)
				}

				r := strings.Split(op.Path, "/")
				collection, recordKey := r[0], r[1]
				op := handlers.OpMeta{
					Repo:       evt.Repo,
					Collection: collection,
					RecordKey:  recordKey,
					Cid:        cid,
				}

				switch record := rec.(type) {
				case *bsky.FeedPost:
					go s.recordHandlers.FeedPostHandler(op, record)
					if err != nil {
						s.logger.Error("FeedPost_RecordHandler", "msg", err)
					}
				}

			}
			return nil
		},
	}

	sched := sequential.NewScheduler("myfirehose", rsc.EventHandler)
	err = events.HandleRepoStream(context.Background(), con, sched, s.logger)
	s.logger.Error("server", "err", err)
	return err
}
