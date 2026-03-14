package main

import (
	"log"

	"github.com/sawada-naoya/mini-redis/internal/server"
)

func main() {
	s := server.New(":6379")

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
