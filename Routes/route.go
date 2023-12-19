package Routes

import (
	"github.com/labstack/echo"
	"net/http"
)

type Routes struct {
}

func (app *Routes) CollectRoutes(e *echo.Echo) {
	appRoutes := e
	product := appRoutes.Group("/product")
	product.GET("/create", func(ctx echo.Context) error {
		data := "create"
		return ctx.String(http.StatusOK, data)
	})

	product.GET("/update", func(ctx echo.Context) error {
		data := "update"
		return ctx.String(http.StatusOK, data)
	})

	product.GET("/delete", func(ctx echo.Context) error {
		data := "delete"
		return ctx.String(http.StatusOK, data)
	})

	product.GET("/list", func(ctx echo.Context) error {
		data := "list"
		return ctx.String(http.StatusOK, data)
	})

	product.GET("/detail", func(ctx echo.Context) error {
		data := "detail"
		return ctx.String(http.StatusOK, data)
	})

	appRoutes.Start(":3000")
}
