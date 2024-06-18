package health

import (
	"micro-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetClient struct {
}

func Get() *GetClient {
	return &GetClient{}
}

func (p *GetClient) Method() string {
	return http.MethodGet
}

func (p *GetClient) Endpoint() string {
	return "/health"
}

// /health
//
//	@Summary		/health エンドポイントハンドラー
//	@Description	/health エンドポイントハンドラー
//	@Accept			json
//	@Produce		json
//	@Tags			health
//	@Success		200	{object}	utils.ResponseData	"result"
//	@Router			/health [get]
func (p *GetClient) Fn(c echo.Context) error {
	return utils.NewResponse(c, http.StatusOK, "OK", nil)
}
