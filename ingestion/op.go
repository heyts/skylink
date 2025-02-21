package ingest

import "github.com/ipfs/go-cid"

type Op struct {
	Repo string
	Path string
	Cid  cid.Cid
}
