package infrastructure

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	entity "../entity"
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

	e.GET("/filter", func(c echo.Context) error {
		fmt.Println("query", c.QueryString())
		ctx := context.Background()
		filter := new(entity.Filter)
		fmt.Println(c.QueryParams())
		// TODO QueryParamsから構造体へのバインディングをやるレシーバ関数を
		// Filter構造体に追加する
		// filter.Keywords = c.QueryParam("keywords")
		filter.Star = entity.Star{Low: c.QueryParam("star[low]"), High: c.QueryParam("star[high]")}
		filter.Language = c.QueryParam("language")
		filter.License = c.QueryParam("license")

		repos, err := repository.Repositories(ctx, filter.BuildQuery())
		if err != nil {
			e.Logger.Fatal(err)
		}
		return c.JSON(http.StatusOK, repos)
	})

	e.Logger.Fatal(e.Start(":1234"))
}

func builtIntoQuery(query string, part string) string {
	if strings.Trim(part, " ") != "" {
		query = query + part
	}
	return query
}
