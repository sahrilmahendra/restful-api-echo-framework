package routes

import (
	"echo-rest/controllers"
	"echo-rest/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is framework echo for golang")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai, middleware.IsAuthenticated)
	e.POST("/pegawai", controllers.StorePegawai, middleware.IsAuthenticated)
	e.PUT("/pegawai", controllers.UpdatePegawai, middleware.IsAuthenticated)
	e.DELETE("/pegawai", controllers.DeletePegawai, middleware.IsAuthenticated)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	e.GET("/test-struct-validation", controllers.TestStructValidation)
	e.GET("/test-variable-validation", controllers.TestVariableValidation)
	return e
}
