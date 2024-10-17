package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func createHTTPServer(req *http.Request) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return rec, c
}

// TestHandleDateTimeCurrent tests if the current time is returned when no date is provided
func TestHandleDateTimeCurrent(t *testing.T) {
	// Initialize Echo instance

	// Create a request with no date parameter
	req := httptest.NewRequest(http.MethodGet, "/api", nil)
	rec, c := createHTTPServer(req)

	if assert.NoError(t, HandleDateTime(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		assert.Contains(t, rec.Body.String(), "unix")
		assert.Contains(t, rec.Body.String(), "utc")
	}
}

// TestHandleDateTimeSpecificDate tests if the handler correctly parses a specific date
func TestHandleDateTimeSpecificDate(t *testing.T) {
	// Create a request with a date parameter
	req := httptest.NewRequest(http.MethodGet, "/api?date=2015-12-25", nil)
	rec, c := createHTTPServer(req)

	if assert.NoError(t, HandleDateTime(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Fri, 25 Dec 2015 00:00:00 UTC")
	}
}

// TestHandleDateTimeInvalidDate tests how the handler handles an invalid date
func TestHandleDateTimeInvalidDate(t *testing.T) {

	// Create a request with an invalid date parameter
	req := httptest.NewRequest(http.MethodGet, "/api?date=invalid-date", nil)
	rec, c := createHTTPServer(req)

	if assert.NoError(t, HandleDateTime(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "Invalid Date")
	}
}
