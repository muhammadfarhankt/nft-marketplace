package middlewaresRepositories

import "github.com/jmoiron/sqlx"

type NMiddlewaresRepository interface {

}

type middlewaresRepository struct {
	db *sqlx.DB
}

func MiddlewaresRepository(db *sqlx.DB) NMiddlewaresRepository {
	return &middlewaresRepository{
		db: db,
	}
}