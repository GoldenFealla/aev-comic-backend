package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"

	"Goldenfealla/aev-comic/config"
	"Goldenfealla/aev-comic/internal/database/postgres"
	"Goldenfealla/aev-comic/internal/repository"
	"Goldenfealla/aev-comic/internal/rest"
	"Goldenfealla/aev-comic/usecase"
)

var (
	e = echo.New()
)

func initMiddleware(cors middleware.CORSConfig) {
	e.Use(middleware.CORSWithConfig(cors))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Request timeout",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			log.Println(c.Path())
		},
		Timeout: 10 * time.Second,
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		CustomTimeFormat: "15:04:05 02/01/2006",
		Format:           "[${time_custom}] ${status} ${method} ${path} ${latency_human} ${error}\n",
		Output:           e.Logger.Output(),
	}))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(20))))
}

func init() {
	cfg := config.Load()

	initMiddleware(cfg.CORS)
	err := postgres.New(cfg.Env.PostgresURI)
	if err != nil {
		log.Fatalf("error init postgres\n error: %v", err)
	}

	comicRepository := repository.NewComicRepository()
	comicUsecase := usecase.NewComicUsecase(comicRepository)
	rest.NewComicHandler(e, comicUsecase)
}

func main() {
	e.Start(":3000")
}
