package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	e.GET("/products", h.FindProduct)
	e.GET("/product/:id", h.GetProduct)
	e.GET("/product-partner/:id", h.FindProductPartner)
	e.POST("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct)))
	e.PATCH("/product/:id", middleware.Auth(h.UpdateProduct))
	e.DELETE("/product/:id", middleware.Auth(h.DeleteProduct))

}
