package main

import (
	"flag"
	"fmt"
	"log"

	"net/http"
	_ "net/http/pprof"

	ingest "github.com/heyts/skylinks"
)

var dsn = flag.String("dsn", "postgres://skylink@localhost/skylink_development?sslmode=disable", "The datasource name to connect to")

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	server := ingest.NewServer(dsn, 10)
	err := server.Start()
	if err != nil {
		fmt.Printf("Server: %v", err)
	}
}
