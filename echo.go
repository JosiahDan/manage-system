package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type DefaultValidatort struct {
	validator *validator.Validate
}

func (v *DefaultValidatort) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// 启动web服务
func runEcho(port int) {
	e := echo.New()
	e.HideBanner = true
	e.Validator = &DefaultValidatort{validator: validator.New()}

	startAt := time.Now()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"startAt": startAt,
			"uptime":  time.Since(startAt).String(),
		})
	})

	setupRoutes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
