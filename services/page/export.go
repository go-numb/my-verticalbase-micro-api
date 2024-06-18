package page

import (
	"fmt"
	"micro-api/utils"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/labstack/echo/v4"
)

type ExportClient struct {
}

func Export() *ExportClient {
	return &ExportClient{}
}

func (p *ExportClient) Method() string {
	return http.MethodGet
}

func (p *ExportClient) Endpoint() string {
	return "/pages/export"
}

// @Router			/pages/export [get]
// @Summary		"Get Pages エンドポイントハンドラー"
// @Description	"ページ情報群を取得するためのエンドポイントハンドラー"
// @Accept			json
// @Produce		json
// @Tags			pages/export
// @Param			limit	query	int	false	"取得するページ数を指定します。	デフォルトは10です。"
// @Success		200	header	file				"取得データはresultに格納されます。"
// @Failure		400	{object}	utils.ResponseData	"idが指定されていない場合に返却されます。"
// @Failure		404	{object}	utils.ResponseData	"ページが見つからない場合に返却されます。"
func (p *ExportClient) Fn(c echo.Context) error {
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

	// make export file
	file := "export.csv"
	// Create a temp file to write CSV data
	tmpFile, err := os.CreateTemp("", file)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	// Write models data to CSV
	w := gocsv.DefaultCSVWriter(tmpFile)
	if err := gocsv.MarshalCSV(&results, w); err != nil {
		return err
	}

	// Close the file so we can serve it
	tmpFile.Close()

	// Download the file
	return c.Attachment(tmpFile.Name(), file)
}
