package infrastructure

import (
	"context"
	"net/http"

	repository "../usecase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.GET("/list", func(c echo.Context) error {
		ctx := context.Background()
		repos, err := repository.Repositories(ctx, "rest language:golang stars:>=200")
		if err != nil {
			e.Logger.Fatal(err)
		}
		return c.JSON(http.StatusOK, repos)
	})

	e.Logger.Fatal(e.Start(":1234"))
}
