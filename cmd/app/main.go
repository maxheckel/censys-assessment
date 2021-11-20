package main

import (
	"github.com/maxheckel/censys-assessment/internal/config"
	"github.com/maxheckel/censys-assessment/internal/server"
)

func main() {
	cfg, err := config.Load("censys")
	if err != nil {
		panic(err)
	}
	srv, err := server.NewServer(cfg)
	if err != nil {
		panic(err)
	}
	srv.Start()
}