package main

import (
	"flag"

	ingest "github.com/heyts/skylinks"
)

var dsn = flag.String("dsn", "file:../sql/skylink.db", "The datasource name to connect to")

func main() {
	flag.Parse()

	server := ingest.NewServer(dsn)
	server.Start()
}
