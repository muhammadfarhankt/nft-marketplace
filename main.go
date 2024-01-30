package main

import (
	"fmt"
	"os"

	"github.com/muhammadfarhankt/nft-marketplace/config"
	"github.com/muhammadfarhankt/nft-marketplace/modules/servers"
	"github.com/muhammadfarhankt/nft-marketplace/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	fmt.Println("Hello, World!")
	config.Init()
	cfg := config.Loadconfig(envPath())
	fmt.Println(cfg.App())
	fmt.Println(cfg.Db())
	fmt.Println(cfg.Jwt())

	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()
}