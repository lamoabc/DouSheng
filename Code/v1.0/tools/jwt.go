package tools

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name"`
	jwt.StandardClaims
}

var myKey = []byte("v1.0")

// GenerateToken
// 生成 token
func GenerateToken(id int64, name string) (string, error) {
	UserClaim := &UserClaims{
		Id:             id,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken
// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}
