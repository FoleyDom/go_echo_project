package controller

import (
	"strconv"

	"github.com/foleydom/go_echo_project/models"
	"github.com/foleydom/go_echo_project/view"
	"github.com/foleydom/go_echo_project/view/layout"
	"github.com/foleydom/go_echo_project/view/user"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct{}

var db = models.ConnectDB()

func (h IndexHandler) Index(c echo.Context) error {
	return render(c, view.Index(), layout.Layout())
}

func (h IndexHandler) GetTodos(c echo.Context) error {
	todos := models.GetTodos(db)
	return c.JSON(200, todos)
}

func (h IndexHandler) GetTodosByChecked(c echo.Context) error {
	checked := true
	todos := models.GetTodosByChecked(db, checked)
	return c.JSON(200, todos)
}

func (h IndexHandler) CreateTodos(c echo.Context) error {
	// id := c.FormValue("id")
	text := c.FormValue("name")
	if text == "" {
		return echo.NewHTTPError(400, "name is required")
	}
	checked := c.FormValue("checked")
	if checked == "" {
		checked = "false"
	}

	todo := models.Todo{
		// ID:      id,
		Text:    text,
		Checked: checked == "true",
	}

	models.CreateTodo(db, &todo)
	if err := db.Error; err != nil {
		return err
	}

	return render(c, user.User(), layout.Layout())
}

func (h IndexHandler) UpdateTodos(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	name := c.FormValue("name")
	if name == "" {
		return echo.NewHTTPError(400, "name is required")
	}
	checked := c.FormValue("checked")
	if checked == "" {
		checked = "false"
	}

	models.UpdateTodos(db, &models.Todo{
		ID:      int(id),
		Text:    name,
		Checked: checked == "true",
	})
	if err := db.Error; err != nil {
		return err
	}

	return c.Redirect(302, "/")
}

func (h IndexHandler) DeleteTodos(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	models.DeleteTodos(db, &models.Todo{ID: int(id)})
	if err := db.Error; err != nil {
		return err
	}

	return c.Redirect(302, "/")
}