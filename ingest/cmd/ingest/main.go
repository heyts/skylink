package main

import (
	ingest "github.com/heyts/skylinks"
)

func main() {
	server := ingest.NewServer()
	server.Start()
}
