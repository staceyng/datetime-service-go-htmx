package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for datetime-service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /api

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/hello-world", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api", HandleDateTime)

	e.Logger.Fatal(e.Start(":1323"))
}

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
	fmt.Println("xxxx ", d)
	dateFormat := "2006-01-02" // YYYY-MM-DD
	parsedTime, err := time.Parse(dateFormat, d)
	fmt.Println("xxxx ", parsedTime, err)

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
