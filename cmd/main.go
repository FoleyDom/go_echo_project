package main

import (
	"github.com/foleydom/go_echo_project/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	UserHandler := handler.UserHandler{}
	Hello_WorldHandler := handler.Hello_WorldHandler{}

	app.GET("/", Hello_WorldHandler.HandleHello_World)
	app.GET("/users", UserHandler.HandleUserShow)
	app.Start(":8080")
}
