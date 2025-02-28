package main

import (
	"flag"

	ingest "github.com/heyts/skylinks"
)

// var dsn = flag.String("dsn", "file:../sql/skylink.db", "The datasource name to connect to")
var dsn = flag.String("dsn", "postgres://skylink@localhost/skylink_development?sslmode=disable", "The datasource name to connect to")

func main() {
	server := ingest.NewServer(dsn)
	server.Start()
}
