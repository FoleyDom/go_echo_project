package main

import (
	"github.com/foleydom/go_echo_project/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	UserHandler := handler.UserHandler{}
	// Hello_WorldHandler := handler.Hello_WorldHandler{}
	IndexHandler := handler.IndexHandler{}

	// app.GET("/", Hello_WorldHandler.HandleHello_World)
	app.GET("/users", UserHandler.HandleUserShow)

	// test route
	app.GET("/test", IndexHandler.Index)

	// Routes
	app.GET("/", IndexHandler.GetTodos)
	app.GET("/checked", IndexHandler.GetTodosByChecked)
	app.POST("/create", IndexHandler.CreateTodos)
	app.POST("/update", IndexHandler.UpdateTodos)
	// app.POST("/delete", IndexHandler.DeleteTodos)

	app.Logger.Fatal(app.Start(":8080"))

	app.Static("/css", "css")
}
