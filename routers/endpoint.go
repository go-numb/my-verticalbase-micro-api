package routers

import (
	"micro-api/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "micro-api/docs"

	sw "github.com/swaggo/echo-swagger"
)

func New(e *echo.Echo) {
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORS(),
	)

	e.GET("/swagger/*", sw.WrapHandler)

	for _, ctrl := range controllers.New() {
		switch ctrl.Method() {
		case http.MethodGet:
			e.GET(ctrl.Endpoint(), ctrl.Fn)
		case http.MethodPost:
			e.POST(ctrl.Endpoint(), ctrl.Fn)
		case http.MethodPut:
			e.PUT(ctrl.Endpoint(), ctrl.Fn)
		case http.MethodDelete:
			e.DELETE(ctrl.Endpoint(), ctrl.Fn)
		}
	}
}
