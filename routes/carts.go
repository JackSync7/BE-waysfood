package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func CartsRoutes(e *echo.Group) {
	CartsRepository := repositories.RepositoryCarts(mysql.DB)
	h := handlers.HandlerCarts(CartsRepository)

	e.GET("/carts", h.FindCarts)
	e.GET("/cart/:id", h.GetCarts)
	e.POST("/cart", middleware.Auth(h.CreateCarts))
	e.DELETE("/cart/:id", middleware.Auth(h.DeleteCarts))

}
