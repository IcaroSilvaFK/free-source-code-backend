package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/IcaroSilvaFK/free-code-source-back/infra/utils"
	"github.com/labstack/echo/v4"
)

type LoginController struct {
}

type LoginControllerInterface interface {
	Login(ctx echo.Context) error
}

func NewLoginController() LoginControllerInterface {
	return &LoginController{}
}

func (lc *LoginController) Login(ctx echo.Context) error {

	code := ctx.Param("code")

	if code == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "code is required",
		})
	}

	fmt.Println(code)

	bt, err := json.Marshal(map[string]string{
		"client_id":     os.Getenv(utils.GITHUB_CLIENT_ID),
		"client_secret": os.Getenv(utils.GITHUB_CLIENT_SECRET),
		"code":          code,
		// "redirect_uri":  "http://localhost:8080/login/github/callback",
	})

	if !errors.Is(err, nil) {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	buff := bytes.NewBuffer(bt)

	c := http.Client{}

	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", buff)

	if !errors.Is(err, nil) {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	r, err := c.Do(req)

	if !errors.Is(err, nil) {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	var authCredentials struct {
		AccessToken string `json:"access_token"`
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	log.Println(string(body))

	if !errors.Is(err, nil) {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	err = json.Unmarshal(body, &authCredentials)

	if !errors.Is(err, nil) {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": authCredentials.AccessToken,
	})

}
