package main

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// echo返回错误json结构（for swagger）
type EchoError struct {
	Message string `json:"message" example:"错误信息"` // 错误信息
} // @name 错误信息

func CheckInput[T any](c echo.Context) (*T, error) {
	var input T
	if err := c.Bind(&input); err != nil {
		return nil, err
	}
	if err := c.Validate(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func JwtEnabled() bool {
	return cfg.JwtSecret != ""
}

func DefaultJwtConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		SigningKey:  []byte(cfg.JwtSecret),
		TokenLookup: "header:Authorization",
		ErrorHandler: func(err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		},
	}
}

func JwtExpireTS() int64 {
	return time.Now().Add(cfg.JwtExpire).Unix()
}

func MakeJwtToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JwtSecret))
}

// OrgFromToken 从token解析出组织
func OrgFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	org := claims["org"].(string)

	return org
}

// 获取token有效期
func GetTokenExp(c echo.Context) int64 {
	return 0
}
