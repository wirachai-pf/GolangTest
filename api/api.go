package api

import (
	"fmt"

	v1 "github.com/wirachai.pf/GolangTest/api/v1"

	"github.com/labstack/echo/v4"
)

func InitWebApiEndpoint(e *echo.Echo) {
	// Web Api
	webApiPrefix := "/api"
	// apply api v1
	applyWebApiV1Endpoint(e, webApiPrefix)
}

func applyWebApiV1Endpoint(e *echo.Echo, apiPrefix string) {
	api := e.Group(fmt.Sprintf("%s/v1", apiPrefix))
	test := api.Group("/test")
	test.GET("/in", v1.GetMockEndpoint)

}
