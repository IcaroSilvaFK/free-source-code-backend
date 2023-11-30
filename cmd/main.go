package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/IcaroSilvaFK/free-code-source-back/cmd/routes"
	"github.com/IcaroSilvaFK/free-code-source-back/infra/utils"
	"github.com/IcaroSilvaFK/free-code-source-back/infra/validators"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//TODO remove comment to enable .env in development
	// if err := godotenv.Load(); !errors.Is(err, nil) {
	// 	log.Fatal("Error loading .env file")
	// }

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	e.Validator = &validators.CustomValidator{
		Validator: validator.New(),
	}

	e.GET("/heath", func(ctx echo.Context) error {

		return ctx.JSON(http.StatusAccepted, echo.Map{
			"message":   "ok",
			"status":    http.StatusAccepted,
			"timestamp": time.Now(),
		})
	})

	r := e.Group("/api")

	routes.NewAppRoutes(r)

	port := os.Getenv(utils.PORT)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
