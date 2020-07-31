package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

//TextIndex validates service Index.
func TestIndex(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, Index(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

//TestPostService validates different inputs and outputs.
func TestPostService(t *testing.T) {
	// Setup
	testsCoupon := make(map[string]float32)

	testsCoupon[`{"amount": 2500, "item_ids": ["MLA710902496", "MLA739047002", "MLA621847666"]}`] = 2500
	testsCoupon[`{"amount": 3500, "item_ids": ["MLA710902496", "MLA739047002", "MLA621847666"]}`] = 3399
	testsCoupon[`{"amount": 5000, "item_ids": ["MLA710902496", "MLA739047002", "MLA621847666"]}`] = 4898
	//Same products.
	testsCoupon[`{"amount": 50000, "item_ids": ["MLA621847666", "MLA621847666", "MLA621847666"]}`] = 2500

	e := echo.New()
	data := &response{}
	for key, value := range testsCoupon {
		req := httptest.NewRequest(http.MethodPost, "/coupon/", strings.NewReader(key))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Assertions
		if assert.NoError(t, postService(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			json.Unmarshal([]byte(strings.Trim(rec.Body.String(), "\n")), data)
			assert.Equal(t, value, data.Total)
		}
	}
}

//TestBadAmount validates if the input parameter 'amount' has a wrong type.
func TestBadAmount(t *testing.T) {
	// Setup
	itemsJSON := `{"amount": "2500", "item_ids": ["MLA710902496"]}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/coupon/", strings.NewReader(itemsJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, postService(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

//TestBadAmount validates if the input parameter 'items' have a wrong type or wrong value.
func TestBadItems(t *testing.T) {
	// Setup
	var itemsJSON [2]string
	itemsJSON[0] = `{"amount": 2500, "item_ids": []}`
	itemsJSON[1] = `{"amount": 10, "item_ids": ["MLA1FAIL", "MLA2FAIL", "MLA3FAIL"]}`
	e := echo.New()

	for _, item := range itemsJSON {
		req := httptest.NewRequest(http.MethodPost, "/coupon/", strings.NewReader(item))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Assertions
		if assert.NoError(t, postService(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	}
}

//TestNotEnoughBuyItem Insufficient amount to buy an item
func TestNotEnoughBuyItem(t *testing.T) {
	// Setup
	itemsJSON := `{"amount": 10, "item_ids": ["MLA710902496", "MLA739047002", "MLA621847666"]}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/coupon/", strings.NewReader(itemsJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, postService(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}
