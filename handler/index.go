package handler

import (
	"github.com/foleydom/go_echo_project/models"
	"github.com/foleydom/go_echo_project/view"
	"github.com/foleydom/go_echo_project/view/layout"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct{}

var db = models.ConnectDB()

func (h IndexHandler) HandleIndex(c echo.Context) error {
	return render(c, view.Index(), layout.Layout())
}

func (h IndexHandler) HandleCreateTodos(c echo.Context) error {
	id := c.FormValue("id")
	text := c.FormValue("text")
	checked := c.FormValue("checked")

	todo := models.Todo{
		ID:      id,
		Text:    text,
		Checked: checked == "on",
	}

	models.CreateTodo(db, &todo)

	return render(c, view.Index(), layout.Layout())
}
