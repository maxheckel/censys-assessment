package main

import (
	"github.com/maxheckel/censys-assessment/internal/config"
	"github.com/maxheckel/censys-assessment/internal/server"
)

func main() {
	srv := server.NewServer(&config.Config{})
	srv.Start()
}