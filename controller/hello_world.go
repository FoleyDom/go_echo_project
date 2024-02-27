package controller

import (
	"github.com/foleydom/go_echo_project/view/layout"
	"github.com/foleydom/go_echo_project/view/user"
	"github.com/labstack/echo/v4"
)

type Hello_WorldHandler struct{}

func (h Hello_WorldHandler) HandleHello_World(c echo.Context) error {
	return render(c, user.Hello_World(), layout.Layout())
}
