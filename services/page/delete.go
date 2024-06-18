package page

import (
	"context"
	"fmt"
	"micro-api/utils"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
)

type DeleteClient struct {
}

func Delete() *DeleteClient {
	return &DeleteClient{}
}

func (p *DeleteClient) Method() string {
	return http.MethodDelete
}

func (p *DeleteClient) Endpoint() string {
	return "/page"
}

// @Router			/page [delete]
// @Summary		"Delete エンドポイントハンドラー"
// @Description	"ページ情報を削除するためのエンドポイントハンドラー"
// @Accept			json
// @Produce		json
// @Tags			page
// @Param			id	query		string	true	"削除するpage idを指定します。"
// @Success		200	{object}	utils.ResponseData		"ステータスコードとメッセージのみ返却されます"
// @Failure		400	{object}	utils.ResponseData		"page bodyが指定されていない場合に返却されます。"
// @Failure		500	{object}	utils.ResponseData		"ページが削除できない場合に返却されます。"
func (p *DeleteClient) Fn(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return utils.NewResponse(c, http.StatusBadRequest, "id is required", nil)
	}

	if err := p.softDelete(c.Request().Context(), id); err != nil {
		return utils.NewResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}

	return utils.NewResponse(c, http.StatusOK, fmt.Sprintf("deleted page, page id: %s", id), nil)
}

func (p *DeleteClient) softDelete(ctx context.Context, id string) error {
	store, err := firestore.NewClient(ctx, os.Getenv("PROJECTID"))
	if err != nil {
		return err
	}
	defer store.Close()

	if _, err = store.Collection("pages").Doc(id).Update(ctx, []firestore.Update{
		{
			Path:  "delete",
			Value: true,
		},
	}); err != nil {
		return err
	}

	return nil
}
