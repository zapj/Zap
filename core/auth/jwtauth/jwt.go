package jwtauth

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zapj/zap/core/global"
)

var SigningKey = []byte("AllYourBase")
var ErrJwtUnknownClaimsType = errors.New("unknown claims type, cannot proceed")

func GenerateAccessToken() (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		Issuer:    "ZAP",
		ID:        "",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access_token, err := token.SignedString(SigningKey)

	if err != nil {
		return "", err
	}

	return access_token, nil
}

func GenerateAccessAndRefreshToken() (gin.H, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		Issuer:    "ZAP",
		ID:        "",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access_token, err := token.SignedString(SigningKey)

	if err != nil {
		return gin.H{}, err
	}
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	refresh_token, err := refreshToken.SignedString([]byte(SigningKey))
	if err != nil {
		return gin.H{}, err
	}
	return gin.H{"access_token": access_token, "refresh_token": refresh_token, "expire_at": time.Minute * 10, "code": 0}, nil
}

func CheckJwtToken(jwtToken string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		global.LOG.Error(err)
		return nil, err
	} else if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok {
		return claims, err
	}
	return nil, ErrJwtUnknownClaimsType
}
