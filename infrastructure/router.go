package infrastructure

import (
	"context"
	"net/http"

	repository "../usecase"
	"github.com/labstack/echo"
)

func Run() {
	e := echo.New()

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
