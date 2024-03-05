package handlers

import (
	"strconv"

	"github.com/foleydom/go_echo_project/models"
	view "github.com/foleydom/go_echo_project/views"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct{}

var db = models.ConnectDB()

func (h IndexHandler) Index(c echo.Context) error {
	return render(c, view.Index([]models.Todo{}))
}

func (h IndexHandler) GetTodos(c echo.Context) error {
	todos := models.GetTodos(db)
	return render(c, view.Index(todos))
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

	return render(c, view.Index([]models.Todo{}))
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
