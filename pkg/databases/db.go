package databases

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/muhammadfarhankt/nft-marketplace/config"
)

func DbConnect(cfg config.IDbConfig) *sqlx.DB {
	db, err := sqlx.Connect("pgx", cfg.Url())
	if err != nil {
		log.Fatalf("connect to db failed: %v\n", err)
	}
	db.DB.SetMaxOpenConns(cfg.MaxOpenConns())
	fmt.Println("DbConnect db:",db)
	return db
}
