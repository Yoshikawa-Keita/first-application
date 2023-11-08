package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type FizzBuzzRequest struct {
	UnixTime int64 `json:"unix_time"`
}

type FizzBuzzResponse struct {
	Result string `json:"result"`
}

func FizzBuzz(c echo.Context) error {
	req := &FizzBuzzRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	var result string
	switch {
	case req.UnixTime%15 == 0:
		result = "Fizz Buzz"
	case req.UnixTime%3 == 0:
		result = "Fizz"
	case req.UnixTime%5 == 0:
		result = "Buzz"
	default:
		result = "Not Fizz Buzz"
	}
	res := &FizzBuzzResponse{Result: result}
	return c.JSON(http.StatusOK, res)
}
