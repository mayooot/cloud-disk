package helper

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func Md5(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

func GenerateToken(id int, identity, name string) (string, error) {
	// id identity name
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
