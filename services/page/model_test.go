package page_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"micro-api/routers"
	"micro-api/services/page"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/joho/godotenv"
)

func TestDo(t *testing.T) {
	godotenv.Load("../../.env")

	// Echoインスタンスを作成
	e := echo.New()
	routers.New(e)

	// レスポンスレコーダーを作成
	rec := httptest.NewRecorder()

	// post as create
	// HTTPリクエストを作成
	id := "6"
	title := "ボブ"
	createPage := page.Page{
		ID:    id,
		Title: title,

		Details: map[string]any{
			"address": "value",
			"tel":     "value2",
			"more": []string{
				"more1",
				"more2",
			},
			"map": map[string]any{
				"lat":  123.456,
				"long": 456.789,
			},
		},

		Categories: []string{"category1", "category2"},
		Tags:       []string{"tag1", "tag2"},

		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	b, err := json.Marshal(createPage)
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/page", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// Routersに登録したハンドラを呼び出す
	e.ServeHTTP(rec, req)
	// ステータスコードのチェック
	fmt.Println("rec: ", rec.Body.String())
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Get
	// HTTPリクエストを作成
	req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/page/%s", id), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// レスポンスレコーダーを作成
	rec = httptest.NewRecorder()

	// Routersに登録したハンドラを呼び出す
	e.ServeHTTP(rec, req)
	// ステータスコードのチェック
	assert.Equal(t, http.StatusOK, rec.Code)

	// レスポンスボディのチェック
	result := echo.Map{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &result))

	fmt.Println("result: ", result["result"])

	assert.Equal(t, title, result["result"].(map[string]any)["title"])

	// Update
	// HTTPリクエストを作成
	putPage := createPage
	putPage.Title = "アリス"
	b, _ = json.Marshal(putPage)
	req = httptest.NewRequest(http.MethodPut, "/page", bytes.NewBuffer(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// レスポンスレコーダーを作成
	rec = httptest.NewRecorder()

	// Routersに登録したハンドラを呼び出す
	e.ServeHTTP(rec, req)
	// ステータスコードのチェック
	assert.Equal(t, http.StatusOK, rec.Code)

	// レスポンスボディのチェック
	result = echo.Map{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &result))

	fmt.Println("result: ", result["result"])

	assert.Equal(t, putPage.Title, result["result"].(map[string]any)["title"])

	// Delete
	// HTTPリクエストを作成
	req = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/page?id=%s", id), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// レスポンスレコーダーを作成
	rec = httptest.NewRecorder()

	// Routersに登録したハンドラを呼び出す
	e.ServeHTTP(rec, req)
	// ステータスコードのチェック
	assert.Equal(t, http.StatusOK, rec.Code)

	// レスポンスボディのチェック
	result = echo.Map{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &result))

	fmt.Println("result: ", result["result"])

	assert.Equal(t, true, result["result"].(map[string]any)["delete"])
}
