package main

import (
	"github.com/foleydom/go_echo_project/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Static("/css", "css")

	UserHandler := handlers.UserHandler{}
	// Hello_WorldHandler := handler.Hello_WorldHandler{}
	IndexHandler := handlers.IndexHandler{}

	// app.GET("/", Hello_WorldHandler.HandleHello_World)
	app.GET("/users", UserHandler.HandleUserShow)

	// test route
	app.GET("/test", IndexHandler.Index)

	// Routes
	app.GET("/", IndexHandler.GetTodos)
	app.GET("/checked", IndexHandler.GetTodosByChecked)
	app.POST("/create", IndexHandler.CreateTodos)
	app.PUT("/update/:id", IndexHandler.UpdateTodos)
	app.POST("/delete/:id", IndexHandler.DeleteTodos)

	app.Logger.Fatal(app.Start(":8080"))

}
