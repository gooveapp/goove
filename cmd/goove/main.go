package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Start(":3000")
	fmt.Println("Goove is running...")

}
