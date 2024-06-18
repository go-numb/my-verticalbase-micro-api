package controllers

import (
	"micro-api/services/health"
	"micro-api/services/page"
	"micro-api/services/skeleton"

	"github.com/labstack/echo/v4"
)

type Controller interface {
	Method() string
	Endpoint() string
	Fn(e echo.Context) error
}

type Client struct {
	Skeleton Controller
	Health   Controller

	Page Controller
}

func New() []Controller {
	return []Controller{
		skeleton.Get(),
		health.Get(),

		page.Get(),
		page.Post(),
		page.Put(),
		page.Delete(),

		page.All(),
		page.Import(),
		page.Export(),
	}
}
