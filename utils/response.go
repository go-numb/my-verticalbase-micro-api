package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	// 必要な型を内包するため、any型を使用
	Result any `json:"result"`
}

func NewResponse(c echo.Context, status int, msg string, result any) error {
	if msg == "" {
		msg = http.StatusText(status)
	}

	return c.JSON(
		status,
		ResponseData{
			Status:  http.StatusText(status),
			Message: msg,
			Result:  result,
		},
	)
}
