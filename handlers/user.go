package handlers

import (
	"github.com/foleydom/go_echo_project/views/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	return render(c, user.User())
}
