package Routes

import (
	"akasia/Controller"
	"github.com/labstack/echo"
)

type Routes struct {
	Controller Controller.Controller
}

func (app *Routes) CollectRoutes(e *echo.Echo) {
	appRoutes := e

	product := appRoutes.Group("/products")
	product.POST("/", app.Controller.CreateProduct)
	product.PUT("/:id", app.Controller.UpdateProduct)
	product.DELETE("/:id", app.Controller.DeleteProduct)
	product.GET("/", app.Controller.ListProduct)
	product.GET("/:id", app.Controller.DetailProduct)

	appRoutes.Start(":3000")
}
