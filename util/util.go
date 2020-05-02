package util

import (
	"net/http"

	"github.com/labstack/echo"
)

//Response struct
type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

//ErrorBody struct
type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

//ErrorResponse struct
type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}

//SendSuccess utility function
func SendSuccess(c echo.Context, data interface{}) error {
	s := Response{}
	s.Status = "success"
	s.Data = data
	return c.JSON(http.StatusOK, s)
}

//SendError utility function
func SendError(c echo.Context, code string, message string, status string) error {
	e := ErrorBody{}
	e.Status = status
	e.Code = code

	s := ErrorResponse{}
	s.Error = e
	return c.JSON(http.StatusInternalServerError, s)

}

//SendData utility function
func SendData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}
