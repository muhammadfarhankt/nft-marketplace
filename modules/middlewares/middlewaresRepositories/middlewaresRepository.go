package middlewaresRepositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/muhammadfarhankt/nft-marketplace/modules/middlewares"
)

type NMiddlewaresRepository interface {
	FindAccessToken(userId, accessToken string) bool
	FindRole() ([]*middlewares.Role, error)
}

type middlewaresRepository struct {
	db *sqlx.DB
}

func MiddlewaresRepository(db *sqlx.DB) NMiddlewaresRepository {
	return &middlewaresRepository{
		db: db,
	}
}

func (m *middlewaresRepository) FindAccessToken(userId, accessToken string) bool {
	query := `
	SELECT
		(CASE WHEN COUNT(*) = 1 THEN true ELSE false END)
	FROM
		"oauth"
	WHERE "user_id" = $1 AND "access_token" = $2;
	`

	var result bool
	if err := m.db.Get(&result, query, userId, accessToken); err != nil {
		return false
	}
	return true
}

func (m *middlewaresRepository) FindRole() ([]*middlewares.Role, error) {
	query := `
	SELECT
		"id",
		"title"
	FROM
		"roles"
	ORDER BY "id" DESC;
	`

	roles := make([]*middlewares.Role, 0)
	if err := m.db.Select(&roles, query); err != nil {
		return nil, fmt.Errorf("roles are empty: %v", err)
	}
	return roles, nil
}
