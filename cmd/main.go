package main

import (
	"flag"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"

	"Goldenfealla/template-go-echo/config"
	"Goldenfealla/template-go-echo/internal/repository"
	"Goldenfealla/template-go-echo/internal/rest"
	"Goldenfealla/template-go-echo/usecase"
)

var (
	isProd = flag.Bool("prod", false, "Enable production mode")
	e      = echo.New()
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
	flag.Parse()

	cfg := config.Load(*isProd)

	initMiddleware(cfg.CORS)

	bookRepository := repository.NewBookRepository()
	bookUsecase := usecase.NewBookUsecase(bookRepository)
	rest.NewBookHandler(e, bookUsecase)
}

func main() {
	if !*isProd {
		log.Println("Running in Development. To run in Prod use \"-prod\"")
	}

	e.Start(":3000")
}
