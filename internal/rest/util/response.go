package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`  // Filled if the response is successful
	Error   interface{} `json:"error,omitempty"` // Filled if an error occurs
}

type ErrorDetail struct {
	Code   int    `json:"code"`
	Detail string `json:"details"` // Use string to hold error message details
}

func SuccessResponse(c echo.Context, message string, data interface{}) error {
	response := ApiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	return c.JSON(http.StatusOK, response)
}

func ErrorResponse(c echo.Context, code int, message string, err error) error {
	// Convert the error to a string for a more friendly API response
	errorDetail := ""
	if err != nil {
		errorDetail = err.Error()
	}

	response := ApiResponse{
		Status:  "error",
		Message: message,
		Error: ErrorDetail{
			Code:   code,
			Detail: errorDetail,
		},
	}
	return c.JSON(code, response)
}
