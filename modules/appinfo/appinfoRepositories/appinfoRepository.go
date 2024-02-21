package appinfoRepositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/muhammadfarhankt/nft-marketplace/modules/appinfo"
)

type IAppinfoRepository interface {
	FindCategory(req *appinfo.CategoryFilter) ([]*appinfo.Category, error)
	InsertCategory(category []*appinfo.Category) error
	DeleteCategory(categoryId string) error
}

type appinfoRepository struct {
	db *sqlx.DB
}

func AppinfoRepository(db *sqlx.DB) IAppinfoRepository {
	return &appinfoRepository{
		db: db,
	}
}

func (r *appinfoRepository) FindCategory(req *appinfo.CategoryFilter) ([]*appinfo.Category, error) {
	query := `
		SELECT
			"id", "title"
		FROM "categories"
	`
	filterValues := make([]any, 0)
	if req.Title != "" {
		query +=
			`WHERE (LOWER("title") LIKE $1)`
		filterValues = append(filterValues, "%"+strings.ToLower(req.Title)+"%")
	}
	query += ";  "
	category := make([]*appinfo.Category, 0)
	if err := r.db.Select(&category, query, filterValues...); err != nil {
		return nil, err
	}
	return category, nil
}

func (r *appinfoRepository) InsertCategory(category []*appinfo.Category) error {
	ctx := context.Background()
	query := `
		INSERT INTO "categories" ("title")
		VALUES
	`
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	valueStack := make([]any, 0)
	for i, cat := range category {
		valueStack = append(valueStack, cat.Title)
		if i != len(category)-1 {
			query += fmt.Sprintf(`($%d),`, i+1)
		} else {
			query += fmt.Sprintf(`($%d)`, i+1)
		}
	}

	query += `
		RETURNING "id";
	`

	rows, err := tx.QueryxContext(ctx, query, valueStack...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error insert category: %w", err)
	}

	var index int
	for rows.Next() {
		if err := rows.Scan(&category[index].Id); err != nil {
			tx.Rollback()
			return fmt.Errorf("error scan category id: %w", err)
		}
		index++
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error commit category: %w", err)
	}
	return nil
}

func (r *appinfoRepository) DeleteCategory(categoryId string) error {
	ctx := context.Background()
	query := `
		DELETE FROM "categories"
		WHERE "id" = $1;`
	_, err := r.db.ExecContext(ctx, query, categoryId)
	if err != nil {
		return fmt.Errorf("error delete category: %w", err)
	}
	return nil
}
