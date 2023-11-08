package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"time"
	"yoshikawa-keita-first-application/app/bff/util"
)

type FizzBuzzRequest struct {
	UnixTime int64 `json:"unix_time"`
}

type FizzBuzzResponse struct {
	Result string `json:"result"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func Hello(c echo.Context) error {
	name := c.QueryParam("name")

	unixTime := time.Now().Unix()
	reqBody, _ := json.Marshal(FizzBuzzRequest{UnixTime: unixTime})
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load configuration")
	}
	res, err := http.Post(config.HTTPServerURL+"/fizzbuzz", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	fmt.Print(res)
	defer res.Body.Close()

	var fizzBuzzRes FizzBuzzResponse
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	json.Unmarshal(body, &fizzBuzzRes)

	helloRes := HelloResponse{
		Message: fmt.Sprintf("Hello, %s. The current time is %d, FizzBuzz is %s", name, unixTime, fizzBuzzRes.Result),
	}
	return c.JSON(http.StatusOK, helloRes)
}
