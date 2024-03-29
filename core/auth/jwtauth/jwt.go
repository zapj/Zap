package jwtauth

import (
	"errors"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zapj/zap/core/global"
)

var ErrJwtUnknownClaimsType = errors.New("unknown claims type, cannot proceed")

type ZapClaims struct {
	Username string `json:"u"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(id, username string) (string, error) {
	claims := ZapClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 120)),
			Issuer:    "ZAP",
			ID:        id,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access_token, err := token.SignedString([]byte(global.SERVER_CONF.SigningKey))

	if err != nil {
		return "", err
	}

	return access_token, nil
}

func GenerateAccessAndRefreshToken() (gin.H, error) {
	claims := &ZapClaims{
		"",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
			Issuer:    "ZAP",
			ID:        "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access_token, err := token.SignedString(global.SERVER_CONF.SigningKey)

	if err != nil {
		return gin.H{}, err
	}
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	refresh_token, err := refreshToken.SignedString([]byte(global.SERVER_CONF.SigningKey))
	if err != nil {
		return gin.H{}, err
	}
	return gin.H{"access_token": access_token, "refresh_token": refresh_token, "expire_at": time.Minute * 10, "code": 0}, nil
}

func CheckJwtToken(jwtToken string) (*ZapClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &ZapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.SERVER_CONF.SigningKey), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		slog.Error("check jwt token", "err", err)
		return nil, err
	} else if claims, ok := token.Claims.(*ZapClaims); ok {
		return claims, err
	}
	return nil, ErrJwtUnknownClaimsType
}
