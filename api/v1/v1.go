package v1

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func GetMockEndpoint(c echo.Context) error {
	fmt.Println("Hello")
	return nil
}
