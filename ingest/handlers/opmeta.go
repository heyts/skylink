package handlers

import "github.com/ipfs/go-cid"

type OpMeta struct {
	Repo       string
	Collection string
	RecordKey  string
	Cid        cid.Cid
}
