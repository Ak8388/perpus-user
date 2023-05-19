package entity

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	Uid string
	Nim string
}

type TokenAkses struct {
	TokenAkses    string
	RefAksesToken string
	ExpAt         int64
	UidTok        string
	UidRefTok     string
}
