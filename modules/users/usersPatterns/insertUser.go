package usersPatterns

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/muhammadfarhankt/nft-marketplace/modules/users"
)

type IInsertUser interface {
	Customer() (IInsertUser, error)
	Admin() (IInsertUser, error)
	Result() (*users.UserPassport, error)
}

type userReq struct {
	id  string
	req *users.UserRegisterReq
	db  *sqlx.DB
}

type customer struct {
	*userReq
}

type admin struct {
	*userReq
}

func InserUser(db *sqlx.DB, req *users.UserRegisterReq, isAdmin bool) IInsertUser {
	if isAdmin {
		return newAdmin(db, req)
	}
	return newCustomer(db, req)
}

func newCustomer(db *sqlx.DB, req *users.UserRegisterReq) IInsertUser {
	return &customer{
		userReq: &userReq{
			req: req,
			db:  db,
		},
	}
}

func newAdmin(db *sqlx.DB, req *users.UserRegisterReq) IInsertUser {
	return &admin{
		userReq: &userReq{
			req: req,
			db:  db,
		},
	}
}

func (u *userReq) Customer() (IInsertUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `
	INSERT INTO "users" 
	(	"email", 
		"password", 
		"username",
		"role_id"
	)
	VALUES 
	($1, $2, $3, 1)
	RETURNING "id";`

	if err := u.db.QueryRowContext(
		ctx,
		query,
		u.req.Email,
		u.req.Password,
		u.req.Username,
	).Scan(&u.id); err != nil {
		switch err.Error() {
		case "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)":
			return nil, fmt.Errorf("email already exists")
		case "ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)":
			return nil, fmt.Errorf("username already exists")
		default:
			return nil, fmt.Errorf("insert user failed: %v", err)
		}
	}
	return u, nil
}

func (u *userReq) Admin() (IInsertUser, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// defer cancel()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `
	INSERT INTO "users" 
	(	"email", 
		"password", 
		"username",
		"role_id"
	)
	VALUES 
	($1, $2, $3, 2)
	RETURNING "id";`

	if err := u.db.QueryRowContext(
		ctx,
		query,
		u.req.Email,
		u.req.Password,
		u.req.Username,
	).Scan(&u.id); err != nil {
		switch err.Error() {
		case "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)":
			return nil, fmt.Errorf("email already exists")
		case "ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)":
			return nil, fmt.Errorf("username already exists")
		default:
			return nil, fmt.Errorf("insert admin failed: %v", err)
		}
	}
	return u, nil
}

func (u *userReq) Result() (*users.UserPassport, error) {
	query := `
	SELECT 
		json_build_object(
			'user', "t",
			'token', NULL
		)
	FROM (
		SELECT
			"us"."id",
			"us"."username",
			"us"."email",
			"us"."role_id"
		FROM "users" "us"
		WHERE "us"."id" = $1
	) AS "t"
	`

	data := make([]byte, 0)
	if err := u.db.Get(&data, query, u.id); err != nil {
		return nil, fmt.Errorf("user get failed: %v", err)
	}

	user := new(users.UserPassport)
	if err := json.Unmarshal(data, &user); err != nil {
		fmt.Println("Data : ", data)
		fmt.Println("*User : ", *user)
		fmt.Println("&User : ", &user)
		fmt.Println("Error : ", err)
		return nil, fmt.Errorf("user unmarshal failed : %v", err)
	}

	return user, nil
}
