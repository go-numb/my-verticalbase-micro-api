package page

import (
	"context"
	"errors"
	"fmt"
	"micro-api/utils"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo/v4"
)

type AllClient struct {
}

func All() *AllClient {
	return &AllClient{}
}

func (p *AllClient) Method() string {
	return http.MethodGet
}

func (p *AllClient) Endpoint() string {
	return "/pages"
}

// @Router			/pages [get]
// @Summary		"Get Pages エンドポイントハンドラー"
// @Description	"ページ情報群を取得するためのエンドポイントハンドラー"
// @Accept			json
// @Produce		json
// @Tags			pages
// @Param			limit	query	int	false	"取得するページ数を指定します。	デフォルトは10です。"
// @Success		200	{object}	[]Page				"取得データはresultに格納されます。"
// @Failure		400	{object}	utils.ResponseData	"idが指定されていない場合に返却されます。"
// @Failure		404	{object}	utils.ResponseData	"ページが見つからない場合に返却されます。"
func (p *AllClient) Fn(c echo.Context) error {
	var (
		limit = 10
		isACS = false
	)

	q := c.QueryParam("limit")
	if q != "" {
		n, err := strconv.Atoi(strings.TrimSpace(q))
		if err != nil {
			return utils.NewResponse(c, http.StatusBadRequest, fmt.Sprintf("invalid limit, error: %v", err), nil)
		}
		limit = n
	}

	pages, err := all(c.Request().Context())
	if err != nil {
		return utils.NewResponse(c, http.StatusNotFound, fmt.Sprintf("pages not found, error: %v", err), nil)
	}

	results := make([]Page, 0)
	for _, page := range pages {
		if page.Delete {
			continue
		}
		results = append(results, page)
	}

	sort.Slice(results, func(i, j int) bool {
		if isACS {
			return results[i].CreatedAt.Before(results[j].CreatedAt)
		}
		return results[i].CreatedAt.After(results[j].CreatedAt)
	})

	if len(results) > limit {
		results = results[:limit]
	}

	return utils.NewResponse(c, http.StatusOK, fmt.Sprintf("pages length: %d", len(results)), results)
}

func all(ctx context.Context) ([]Page, error) {
	store, err := firestore.NewClient(ctx, os.Getenv("PROJECTID"))
	if err != nil {
		return nil, err
	}
	defer store.Close()

	docs, err := store.Collection("pages").Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var (
		errs  = make([]error, 0)
		pages = make([]Page, 0)
	)
	for _, doc := range docs {
		page := Page{}
		if err := doc.DataTo(&page); err != nil {
			errs = append(errs, err)

			continue
		}
		pages = append(pages, page)
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return pages, nil
}
