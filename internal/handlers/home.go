package handlers

import (
	"github.com/gooveapp/goove/view/home"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
}

func (h HomeHandler) HandleHome(c echo.Context) error {
	return render(c, home.Home())
}
