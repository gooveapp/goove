package handlers

import (
	"github.com/gooveapp/goove/view/record"
	"github.com/labstack/echo/v4"
)

type RecordHandler struct {
}

func (r RecordHandler) HandleRecords(c echo.Context) error {
	return render(c, record.RecordHome())
}

func (r RecordHandler) HandleRecordCreatePage(c echo.Context) error {
	return render(c, record.RecordCreatePage())
}
