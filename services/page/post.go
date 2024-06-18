package page

import (
	"context"
	"fmt"
	"micro-api/utils"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
)

type PostClient struct {
}

func Post() *PostClient {
	return &PostClient{}
}

func (p *PostClient) Method() string {
	return http.MethodPost
}

func (p *PostClient) Endpoint() string {
	return "/page"
}

//	@Router			/page [post]
//	@Summary		"Post エンドポイントハンドラー"
//	@Description	"ページ情報を保存するためのエンドポイントハンドラー"
//	@Accept			json
//	@Produce		json
//	@Tags			page
//	@Param			page	body		object				true	"保存するpage情報を指定します。"
//	@Success		201		{object}	Page				"保存データはresultに格納され、返却されます"
//	@Failure		400		{object}	utils.ResponseData	"page bodyが指定されていない場合に返却されます。"
//	@Failure		500		{object}	utils.ResponseData	"ページが保存できない場合に返却されます。"
func (p *PostClient) Fn(c echo.Context) error {
	page := &Page{}
	if err := c.Bind(page); err != nil {
		return utils.NewResponse(c, http.StatusBadRequest, "Invalid request", nil)
	}

	if page.ID == "" {
		page.ID = xid.New().String()
	}

	if err := p.create(c.Request().Context(), page); err != nil {
		return utils.NewResponse(c, http.StatusInternalServerError, "Internal server error", nil)
	}

	return utils.NewResponse(c, http.StatusCreated, fmt.Sprintf("created page, page id: %s", page.ID), page)
}

func (p *PostClient) create(ctx context.Context, page *Page) error {
	store, err := firestore.NewClient(ctx, os.Getenv("PROJECTID"))
	if err != nil {
		return err
	}
	defer store.Close()

	// 新規作成
	// ID重複があれば、エラーを返す
	if _, err = store.Collection("pages").Doc(page.ID).Create(ctx, page); err != nil {
		return err
	}

	return nil
}
