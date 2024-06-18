package page

import (
	"context"
	"fmt"
	"micro-api/utils"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
)

type PutClient struct {
}

func Put() *PutClient {
	return &PutClient{}
}

func (p *PutClient) Method() string {
	return http.MethodPut
}

func (p *PutClient) Endpoint() string {
	return "/page"
}

//	@Router			/page [put]
//	@Summary		"Put エンドポイントハンドラー"
//	@Description	"ページ情報を更新するためのエンドポイントハンドラー"
//	@Accept			json
//	@Produce		json
//	@Tags			page
//	@Param			page	body		object				true	"更新するpage情報を指定します。"
//	@Success		200		{object}	Page				"更新データはresultに格納され、返却されます"
//	@Failure		400		{object}	utils.ResponseData	"page bodyが指定されていない場合に返却されます。"
//	@Failure		500		{object}	utils.ResponseData	"ページが更新できない場合に返却されます。"
func (p *PutClient) Fn(c echo.Context) error {
	page := &Page{}
	if err := c.Bind(page); err != nil {
		return utils.NewResponse(c, http.StatusBadRequest, "Invalid request", nil)
	}

	if err := p.put(c.Request().Context(), page); err != nil {
		return utils.NewResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}

	return utils.NewResponse(c, http.StatusOK, fmt.Sprintf("updated page, page id: %s", page.ID), page)
}

func (p *PutClient) put(ctx context.Context, page *Page) error {
	store, err := firestore.NewClient(ctx, os.Getenv("PROJECTID"))
	if err != nil {
		return err
	}
	defer store.Close()

	if _, err = store.Collection("pages").Doc(page.ID).Update(ctx, []firestore.Update{
		{
			Path:  "title",
			Value: page.Title,
		},
		{
			Path:  "slug",
			Value: page.Slug,
		},
		{
			Path:  "description",
			Value: page.Description,
		},
		{
			Path:  "details",
			Value: page.Details,
		},
		{
			Path:  "categories",
			Value: page.Categories,
		},
		{
			Path:  "tags",
			Value: page.Tags,
		},
		{
			Path:  "updated_at",
			Value: time.Now(),
		},
	}); err != nil {
		return err
	}

	return nil
}
