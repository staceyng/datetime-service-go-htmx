# datetime-service-go-htmx

Simple datetime microservice written in Go + HTMX.

Task adapted from timstamp microservice from https://www.freecodecamp.org/learn/back-end-development-and-apis/back-end-development-and-apis-projects/timestamp-microservice

## Libraries used

Go Backend

1. [swaggo/swag](https://github.com/swaggo/swag) - to automatically generate swagger file definitions
2. [echo](https://github.com/labstack/echo) - for ease of starting a REST API server

## Testing

Backend testing

1. Run server with command `go run main.go`
2. In another terminal execute curl commands
   - Test valid date eg. 2024-11-16 with `curl "http://localhost:1323/api?date=2024-11-16"`
   - Test valid timestamp eg. 1451001600000 with `curl "http://localhost:1323/api?date=1451001600000"`
   - Test empty date, should return current time with `curl "http://localhost:1323/api"`
