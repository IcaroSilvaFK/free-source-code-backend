package validators

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(input interface{}) error {

	if err := cv.Validator.Struct(input); !errors.Is(err, nil) {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
