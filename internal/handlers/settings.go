package handlers

import (
	"github.com/gooveapp/goove/view/settings"
	"github.com/labstack/echo/v4"
)

type SettingsHandler struct {
}

func (s SettingsHandler) HandleSettings(c echo.Context) error {
	return render(c, settings.Settings())
}
