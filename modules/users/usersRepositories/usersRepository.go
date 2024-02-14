package usersRepositories

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/muhammadfarhankt/nft-marketplace/modules/users"
	"github.com/muhammadfarhankt/nft-marketplace/modules/users/usersPatterns"
)

type IUsersRepository interface {
	InsertUser(req *users.UserRegisterReq, isAdmin bool) (*users.UserPassport, error)
	FindOneUserByUsername(username string) (*users.UserCredentialCheck, error)
	InsertOAuth(req *users.UserPassport) error
	FindOneOAuth(refreshToken string) (*users.Oauth, error)
	UpdateOneOAuth(req *users.UserToken) error
	GetProfile(userId string) (*users.User, error)
	DeleteOauth(oauthId string) error
}

type usersRepository struct {
	db *sqlx.DB
}

func UsersRepository(db *sqlx.DB) IUsersRepository {
	return &usersRepository{
		db: db,
	}
}

func (u *usersRepository) InsertUser(req *users.UserRegisterReq, isAdmin bool) (*users.UserPassport, error) {
	result := usersPatterns.InserUser(u.db, req, isAdmin)
	var err error
	if isAdmin {
		result, err = result.Admin()
		if err != nil {
			return nil, err
		}
	} else {
		result, err = result.Customer()
		if err != nil {
			return nil, err
		}
	}
	//Getting result after insertion
	user, err := result.Result()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *usersRepository) FindOneUserByUsername(username string) (*users.UserCredentialCheck, error) {
	query := `
		SELECT "id", "username", "email", "password", "role_id"
		FROM "users"
		WHERE username = $1;
	`
	user := new(users.UserCredentialCheck)
	if err := r.db.Get(user, query, username); err != nil {
		return nil, fmt.Errorf("user not found w/ username %v \n error: %v", username, err)
	}

	return user, nil

	//return usersPatterns.
	//return usersPatterns.FindOneUserByUsername(r.db, username)
}

func (r *usersRepository) InsertOAuth(req *users.UserPassport) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	query := `
		INSERT INTO "oauth"
		("user_id", "access_token", "refresh_token")
		VALUES
		($1, $2, $3)
		RETURNING "id";
	`
	err := r.db.QueryRowContext(ctx, query, req.User.Id, req.Token.AccessToken, req.Token.RefreshToken).Scan(&req.Token.Id)
	if err != nil {
		return fmt.Errorf("inserting oauth user token failed: %v", err)
	}
	return nil
}

func (r *usersRepository) FindOneOAuth(refreshToken string) (*users.Oauth, error) {
	query := `
		SELECT "id", "user_id"
		FROM "oauth"
		WHERE "refresh_token" = $1;
	`
	oauth := new(users.Oauth)
	if err := r.db.Get(oauth, query, refreshToken); err != nil {
		return nil, fmt.Errorf("oauth not found. Error : %v", err)
	}
	return oauth, nil
}

func (r *usersRepository) UpdateOneOAuth(req *users.UserToken) error {
	query := `
		UPDATE "oauth"
		SET "access_token" = :access_token,
			"refresh_token" = :refresh_token
		WHERE "id" = :id
	;`
	if _, err := r.db.NamedExecContext(context.Background(), query, req); err != nil {
		return fmt.Errorf("updating oauth fialed : %v", err)
	}
	return nil
}

func (r *usersRepository) GetProfile(userId string) (*users.User, error) {
	query := `
	SELECT
		"id",
		"email",
		"username",
		"role_id"
	FROM "users"
	WHERE "id" = $1;`

	profile := new(users.User)
	if err := r.db.Get(profile, query, userId); err != nil {
		return nil, fmt.Errorf("get user failed: %v", err)
	}
	return profile, nil
}

func (r *usersRepository) DeleteOauth(oauthId string) error {
	query := `
		DELETE FROM "oauth"
		WHERE "id" = $1;
	`
	fmt.Println("oauthId : ", oauthId)
	if _, err := r.db.ExecContext(context.Background(), query, oauthId); err != nil {
		return fmt.Errorf("deleting oauth failed : %v", err)
	}
	return nil
}
