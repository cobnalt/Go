package main

import (
	"github.com/cobnalt/Go/internal/config"
	"github.com/cobnalt/Go/internal/database"
	"github.com/cobnalt/Go/internal/router"
	"github.com/cobnalt/Go/internal/service"
	_ "github.com/lib/pq"
)

type product struct {
	id   int
	name string
}

func main() {
	cfg, err := config.ReadConfig("configs/config.toml")
	if err != nil {
		panic(err)
	}

	db, err := database.New(cfg.Database)
	if err != nil {
		panic(err)
	}

	service, err := service.New(db)
	if err != nil {
		panic(err)
	}

	server := router.New(*service)
	server.Run()
}
