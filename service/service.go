package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type DateResponse struct {
	Unix int64  `json:"unix"`
	UTC  string `json:"utc"`
}

// ShowAccount godoc
// @Summary      Show a utc date and epoch timestamp
// @Description  Get a UTC date or timestamp from a given date
// @Tags         datetime
// @Accept       json
// @Produce      json
// @Param        date  path  string  true  "string"
// @Success      200  {object}  DateResponse
// @Failure      400  {object}  ErrorResponse
// @Router       /{date} [get]
func HandleDateTime(c echo.Context) error {
	date := c.QueryParam("date")

	var t time.Time
	var err error

	if date == "" {
		t = time.Now()
	} else {
		t, err = validateDateTime(date)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid Date"})
		}
	}

	return c.JSON(http.StatusOK, DateResponse{Unix: t.UnixMilli(), UTC: t.UTC().Format(time.RFC1123)})
}

func validateDateTime(d string) (time.Time, error) {
	dateFormat := "2006-01-02" // YYYY-MM-DD
	parsedTime, err := time.Parse(dateFormat, d)

	if err == nil {
		return parsedTime, nil
	}

	// Try to parse as an epoch timestamp (in milliseconds)
	epochTime, err := strconv.ParseInt(d, 10, 64)
	if err == nil {
		// Convert milliseconds to seconds for time.Unix
		return time.Unix(0, epochTime*int64(time.Millisecond)), nil
	}

	return time.Time{}, fmt.Errorf("invalid input: not a valid date or epoch time")
}
