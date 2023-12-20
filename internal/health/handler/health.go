package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheckResponse is the response body of health check endpoint
type HealthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// HealthCheck is the handler of health check endpoint
func HealthCheck(c echo.Context) error {
	// TODO: Implement health check logic here
	response := HealthCheckResponse{
		Status:  "OK",
		Message: "Server is runing",
	}

	return c.JSON(http.StatusOK, response)
}
