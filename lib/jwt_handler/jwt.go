package jwthandler

import (
	"errors"
	"go-boilerplate/bootstrap"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyCustomClaims struct {
	UUID				string		`json:"uuid"`
	Name				string		`json:"name"`
	Role				string		`json:"role"`
	jwt.RegisteredClaims
}

var mySigningKey = []byte(bootstrap.App.Config.GetString("jwt.secret"))
var lifetime = bootstrap.App.Config.GetInt("jwt.lifetime")

func GenerateAllTokens(uuid string, name string,  role string) (string, string, error){
	tokenLifetime := time.Duration(lifetime)

	// Create the claims
	claims := &MyCustomClaims{
		UUID: uuid,
		Name: name,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			Issuer:    "Authenticator",
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(tokenLifetime * time.Minute)),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		},
	}
	// create refresh token
	refreshClaims := &MyCustomClaims{
		UUID: uuid,
		Name: name,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Authenticator",
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().AddDate(1, 0, 0)),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		},
	}

	token, err  := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySigningKey)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(mySigningKey)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, err
}

func ValidateToken(signedToken string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok:= token.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("Token is invalid")
	}

	// if claims.ExpiresAt.UTC().Unix() < time.Now().UTC().Unix(){
	// 	return nil, errors.New("Token is expired")
	// }

	return claims, nil
}
