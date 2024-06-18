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

type GetClient struct {
}

func Get() *GetClient {
	return &GetClient{}
}

func (p *GetClient) Method() string {
	return http.MethodGet
}

func (p *GetClient) Endpoint() string {
	return "/page/:id"
}

//	@Router			/page/{id} [get]
//	@Summary		"Get エンドポイントハンドラー"
//	@Description	"ページ情報を取得するためのエンドポイントハンドラー"
//	@Accept			json
//	@Produce		json
//	@Tags			page
//	@Param			id	path		string				true	"取得するpage idを指定します。"
//	@Success		200	{object}	Page				"取得データはresultに格納されます。"
//	@Failure		400	{object}	utils.ResponseData	"idが指定されていない場合に返却されます。"
//	@Failure		404	{object}	utils.ResponseData	"ページが見つからない場合に返却されます。"
func (p *GetClient) Fn(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.NewResponse(c, http.StatusBadRequest, "id is required", nil)
	}

	page, err := p.get(c.Request().Context(), id)
	if err != nil {
		return utils.NewResponse(c, http.StatusNotFound, fmt.Sprintf("page not found, page id: %s, error: %v", id, err), nil)
	}

	if page.Delete {
		return utils.NewResponse(c, http.StatusNotFound, fmt.Sprintf("page not found, page id: %s", id), nil)
	}

	return utils.NewResponse(c, http.StatusOK, "", page)
}

func (p *GetClient) get(ctx context.Context, id string) (*Page, error) {
	store, err := firestore.NewClient(ctx, os.Getenv("PROJECTID"))
	if err != nil {
		return nil, err
	}
	defer store.Close()

	doc, err := store.Collection("pages").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	page := &Page{}
	if err := doc.DataTo(page); err != nil {
		return nil, err
	}

	return page, nil
}
