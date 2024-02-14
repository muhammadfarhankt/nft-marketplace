package nftauth

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammadfarhankt/nft-marketplace/config"
	"github.com/muhammadfarhankt/nft-marketplace/modules/users"
)

type TokenType string

const (
	Access  TokenType = "access"
	Refresh TokenType = "refresh"
	Admin   TokenType = "admin"
	ApiKey  TokenType = "apikey"
)

type nftAuth struct {
	mapClaims *nftMapClaims
	cfg       config.IJwtConfig
}

type INftAuth interface {
	//NewAuth(tokenType TokenType, cfg config.IJwtConfig, claims *users.UserClaims) (nftAuth, error)
	SignToken() string
}

type nftMapClaims struct {
	Claims *users.UserClaims `json:"claims"`
	jwt.RegisteredClaims
}

func jwtTimeDurationCalc(t int) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(time.Duration(int64(t) * int64(math.Pow10(9)))))
}

func jwtTimeRepeatAdapter(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Unix(t, 0))
}

func (n *nftAuth) SignToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, n.mapClaims)
	tokenString, _ := token.SignedString([]byte(n.cfg.SecretKey()))
	return tokenString
	// if err != nil {
	// 	return ""
	// }
	// return tokenString
}

func RepeatToken(cfg config.IJwtConfig, claims *users.UserClaims, exp int64) string {
	obj := &nftAuth{
		cfg: cfg,
		mapClaims: &nftMapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "nft-marketplace",
				Subject:   "refresh-token",
				Audience:  []string{"customer", "admin"},
				ExpiresAt: jwtTimeRepeatAdapter(exp),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
	return obj.SignToken()
}

func ParseToken(cfg config.IJwtConfig, tokenString string) (*nftMapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &nftMapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method ")
		}
		return cfg.SecretKey(), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, fmt.Errorf("invalid token format")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("token expired")
		} else {
			return nil, fmt.Errorf("parse token failed: %v", err)
		}
		// else if errors.Is(err, jwt.ErrSignatureInvalid) {
		// 	return nil, fmt.Errorf("invalid token signature")
		// }
	}

	if claims, ok := token.Claims.(*nftMapClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid claim type %v", err)
	}
}

func NewAuth(tokenType TokenType, cfg config.IJwtConfig, claims *users.UserClaims) (INftAuth, error) {
	switch tokenType {
	case Access:
		return newAccessToken(cfg, claims), nil
	case Refresh:
		return newRefreshToken(cfg, claims), nil
	default:
		return nil, fmt.Errorf("invalid token type")
	}

	// return &nftAuth{
	// 	mapClaims: &nftMapClaims{
	// 		Claims: &users.UserClaims{
	// 			Type: string(tokenType),
	// 		},
	// 	},
	// 	cfg: cfg,
	// },
	// nil,
}

func newAccessToken(cfg config.IJwtConfig, claims *users.UserClaims) INftAuth {
	return &nftAuth{
		cfg: cfg,
		mapClaims: &nftMapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "nft-marketplace",
				Subject:   "access-token",
				Audience:  []string{"customer", "admin"},
				ExpiresAt: jwtTimeDurationCalc(cfg.AccessExpiresAt()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
}

func newRefreshToken(cfg config.IJwtConfig, claims *users.UserClaims) INftAuth {
	return &nftAuth{
		cfg: cfg,
		mapClaims: &nftMapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "nft-marketplace",
				Subject:   "refresh-token",
				Audience:  []string{"customer", "admin"},
				ExpiresAt: jwtTimeDurationCalc(cfg.RefreshExpiresAt()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
}
