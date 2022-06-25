package main

import (
	"net/http"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/artempopkov/infobase-golang/internal/handler"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
  e.Use(middleware.Recover())


	helloHandler := handler.NewHelloHandler()
  // Routes
  e.GET("/", helloHandler.Get)


  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}
