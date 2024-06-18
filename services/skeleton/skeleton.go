package skeleton

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
	return "/skeleton"
}

// /skeleton
//
//	@Summary		/skeleton エンドポイントハンドラー
//	@Description	/skeleton エンドポイントハンドラー
//	@Accept			json
//	@Produce		json
//	@Tags			skeleton
//	@Success		200	{object}	utils.ResponseData	"result"
//	@Router			/skeleton [get]
func (p *GetClient) Fn(c echo.Context) error {
	return utils.NewResponse(c, http.StatusOK, "skeleton endpoint functions", nil)
}
