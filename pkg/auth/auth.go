package auth

import "github.com/dgrijalva/jwt-go"

const JWT_SECRET = "fmfwamk53m43amfacvmtmboqpweokrttk"

type JWTClaims struct {
	UserId   uint   `json: "userid"`
	Username string `json: "username"`
	jwt.StandardClaims
}
