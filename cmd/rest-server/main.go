package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/personal-golang/ad-rest-api/pkg/clients/ad"
	"github.com/personal-golang/ad-rest-api/pkg/components/ad"
	"net/http"
)

func main() {
	adclient.InitConn()
	defer adclient.CloseConn()
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", helloHandler)
	app.POST("/group/:group_name", ad.NewGroupHandler)

	app.Logger.Fatal(app.Start(":1323"))
}

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
