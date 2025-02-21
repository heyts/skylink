package ingest

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/events"
	"github.com/bluesky-social/indigo/events/schedulers/sequential"
	"github.com/bluesky-social/indigo/repo"
	"github.com/gorilla/websocket"
	"github.com/heyts/skylinks/models"
	"github.com/heyts/skylinks/utils"
)

var resolverPolicy = map[string][]string{
	"youtube.com": {"v"},
}

type Server struct {
	logger         *slog.Logger
	wsEndpoint     string
	dsn            string
	db             *sql.DB
	domainResolver *utils.DomainResolver
}

func NewServer(dsn *string) *Server {
	db, err := sql.Open("sqlite3", *dsn)
	if err != nil {
		log.Fatal(err)
	}

	s := &Server{
		logger:         slog.New(slog.NewTextHandler(os.Stderr, nil)),
		wsEndpoint:     "wss://bsky.network/xrpc/com.atproto.sync.subscribeRepos",
		dsn:            *dsn,
		db:             db,
		domainResolver: utils.NewDomainResolver(resolverPolicy),
	}

	return s
}

func (s *Server) Start() error {
	con, _, err := websocket.DefaultDialer.Dial(s.wsEndpoint, http.Header{})
	if err != nil {
		s.logger.Error(fmt.Sprintf("Could not connect to %q, aborting.\n", s.wsEndpoint))
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
				op := OpMeta{evt.Repo, collection, recordKey, cid}

				switch record := rec.(type) {
				case *bsky.FeedPost:
					err := s.HandleFeedPost(op, record)
					if err != nil {
						s.logger.Error("HandleFeedPost", "msg", err)
					}
				}

			}
			return nil
		},
	}

	sched := sequential.NewScheduler("myfirehose", rsc.EventHandler)
	err = events.HandleRepoStream(context.Background(), con, sched, s.logger)
	return err
}

func (s *Server) HandleFeedPost(op OpMeta, record *bsky.FeedPost) error {
	actor, err := s.resolveUserIdentity(op.Repo)
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
					url, err := s.domainResolver.Resolve(ft.RichtextFacet_Link.Uri)
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

	if len(model.URI) > 0 {
		s.logger.Info("record", "val", model)
	}
	return nil
}

func (*Server) resolveUserIdentity(did string) (*bsky.ActorDefs_ProfileView, error) {
	endpoint := fmt.Sprintf("https://public.api.bsky.app/xrpc/app.bsky.actor.getProfile?actor=%s", did)
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	actor := bsky.ActorDefs_ProfileView{}

	err = json.Unmarshal(body, &actor)
	if err != nil {
		return nil, err
	}

	return &actor, nil
}
