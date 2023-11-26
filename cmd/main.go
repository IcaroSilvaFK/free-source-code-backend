package main

import (
	"errors"
	"log"

	"github.com/IcaroSilvaFK/free-code-source-back/cmd/routes"
	"github.com/IcaroSilvaFK/free-code-source-back/infra/validators"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	if err := godotenv.Load(); !errors.Is(err, nil) {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	e.Validator = &validators.CustomValidator{
		Validator: validator.New(),
	}

	r := e.Group("/api")

	routes.NewAppRoutes(r)

	e.Logger.Fatal(e.Start(":8080"))
}
