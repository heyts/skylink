package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"net/http"
	_ "net/http/pprof"

	ingest "github.com/heyts/skylinks"
)

var dsn = flag.String("dsn", os.Getenv("DATABASE_URL"), "The datasource name to connect to")

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
