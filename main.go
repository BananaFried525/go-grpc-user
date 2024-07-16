package main

import (
	"bananaf/wallet/user/boots"
	"bananaf/wallet/user/config"
)

func main() {
	config := config.Config{}
	con, _ := config.NewConfig()

	gs := boots.NewGrpcServer(con)
	gs.Start()
}
