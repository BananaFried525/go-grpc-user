package main

import (
	"github.com/bananafried525/wallet/user/config"
	"github.com/bananafried525/wallet/user/database"
	"github.com/bananafried525/wallet/user/server"
)

func main() {
	config := config.Config{}
	cfg, _ := config.New()

	db := database.New(cfg)
	gs := server.New(cfg, db)

	gs.Start()
}
