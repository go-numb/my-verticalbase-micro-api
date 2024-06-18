package page

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"micro-api/utils"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"

	"github.com/gocarina/gocsv"
)

type ImportClient struct {
}

func Import() *ImportClient {
	return &ImportClient{}
}

func (p *ImportClient) Method() string {
	return http.MethodPost
}

func (p *ImportClient) Endpoint() string {
	return "/pages/import"
}

// @Router			/pages/import [post]
// @Summary		"Post エンドポイントハンドラー"
// @Description	"ページ情報を一括保存するためのエンドポイントハンドラー"
// @Accept			multipart/form-data
// @Tags			pages/import
// @Param			file	formData		file				true	"保存するpage情報群が記載されたCSVファイルを指定します。"
// @Success		201		{object}	[]Page				"保存データはresultに格納され、返却されます"
// @Failure		400		{object}	utils.ResponseData	"page bodyが指定されていない場合に返却されます。"
// @Failure		500		{object}	utils.ResponseData	"ページが保存できない場合に返却されます。"
func (p *ImportClient) Fn(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return utils.NewResponse(c, http.StatusBadRequest, "Invalid request", nil)
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	buffer := make([]byte, file.Size)
	if _, err = src.Read(buffer); err != nil {
		return err
	}

	fmt.Println("buffer: ", string(buffer))

	var s []*PageOTC
	if err := gocsv.Unmarshal(bytes.NewBuffer(buffer), &s); err != nil {
		return utils.NewResponse(c, http.StatusInternalServerError, fmt.Sprintf("Internal server error: %v", err), nil)
	}

	pages := make([]*Page, 0)
	for _, row := range s {
		page := &Page{
			ID:          row.ID,
			Title:       row.Title,
			Slug:        row.Slug,
			Description: row.Description,

			// 以下、省略
			Details: map[string]any{
				// 省略
			},

			Categories: utils.Split(row.Categories),
			Tags:       utils.Split(row.Tags),

			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		}

		pages = append(pages, page)
	}

	if len(pages) == 0 {
		return utils.NewResponse(c, http.StatusBadRequest, "Invalid request, has not length from import file", nil)
	}

	if err := p.registor(c.Request().Context(), pages); err != nil {
		return utils.NewResponse(c, http.StatusInternalServerError, fmt.Sprintf("Internal server error: %v", err.Error()), nil)
	}

	return utils.NewResponse(c, http.StatusCreated, fmt.Sprintf("registored pages, page length: %d", len(pages)), pages)
}

func (p *ImportClient) registor(ctx context.Context, pages []*Page) error {
	store, err := firestore.NewClient(ctx, os.Getenv("PROJECTID"))
	if err != nil {
		return err
	}
	defer store.Close()

	// 新規作成
	// ID重複があれば、エラーを返す
	errs := make([]error, 0)
	for _, page := range pages {
		if _, err = store.Collection("pages").Doc(page.ID).Create(ctx, page); err != nil {
			errs = append(errs, err)
			continue
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
