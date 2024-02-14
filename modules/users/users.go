package users

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	// Password string `db:"password" json:"password"`
	RoleId int `db:"role_id" json:"role_id"`
}

type UserToken struct {
	Id           string `db:"id" json:"id"`
	AccessToken  string `db:"access_token" json:"access_token"`
	RefreshToken string `db:"refresh_token" json:"refresh_token"`
}

type UserPassport struct {
	User  *User      `json:"user"`
	Token *UserToken `json:"token"`
}

type UserRegisterReq struct {
	Username string `db:"username" json:"username" form:"username"`
	Email    string `db:"email" json:"email" form:"email"`
	Password string `db:"password" json:"password" form:"password"`
}

type UserCredential struct {
	//Email   string `db:"email" json:"email" form:"email"`
	Username string `db:"username" json:"username" form:"username"`
	Password string `db:"password" json:"password" form:"password"`
}

type UserCredentialCheck struct {
	Id       string `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
	RoleId   int    `db:"role_id"`
}

type UserClaims struct {
	Id     string `json:"id" db:"id"`
	RoleId int    `json:"role_id" db:"role_id"`
}

type UserRefreshCredential struct {
	RefreshToken string `json:"refresh_token" db:"refresh_token" form:"refresh_token"`
}

type Oauth struct {
	Id     string `db:"id" json:"id"`
	UserId string `db:"user_id" json:"user_id"`
}

type UserRemoveCredential struct {
	OauthId string `json:"oauth_id" db:"id" form:"oauth_id"`
}

func (obj *UserRegisterReq) BcryptHashing() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(obj.Password), 10)
	if err != nil {
		return fmt.Errorf("password hashing failed: %v", err)
	}
	obj.Password = string(hashedPassword)
	//fmt.Println("Obj : ", obj)
	return nil
}

func (obj *UserRegisterReq) IsEmail() bool {
	match, err := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, obj.Email)
	if err != nil {
		return false
	}
	//fmt.Println("Email : ", obj.Email)
	//fmt.Println("error : ", err)
	return match
}
