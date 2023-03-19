package main

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"k8s.io/klog/v2"
)

func setupRoutes(e *echo.Echo) {
	var jwtConfig middleware.JWTConfig

	admin := e.Group("/admin")

	if JwtEnabled() {
		jwtConfig = DefaultJwtConfig()
		admin.Use(middleware.JWTWithConfig(jwtConfig))
	}

	e.POST("/login", Login)

}

func Login(c echo.Context) error {
	var loginInfo LoginInfo
	if err := c.Bind(&loginInfo); err != nil {
		return c.JSON(http.StatusBadRequest, "请求参数错误:"+err.Error())
	}

	if loginInfo.Password == "password" && loginInfo.Username == "admin" {
		token, err := MakeJwtToken(jwt.MapClaims{
			"username": "admin",
			"isAdmin":  true,
		})
		if err != nil {
			klog.Fatal("生成Token失败:", err)
		}
		return c.JSON(http.StatusOK, token)
	}
	return c.JSON(http.StatusForbidden, "账号密码错误")
}
