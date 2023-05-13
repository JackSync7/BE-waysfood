package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func OrderRoutes(e *echo.Group) {
	OrderRepository := repositories.RepositoryOrder(mysql.DB)
	h := handlers.HandlerOrder(OrderRepository)

	e.GET("/orders", h.FindOrder)
	e.GET("/order/:id", h.GetOrder)
	e.GET("/order-product/:id", middleware.Auth(h.GetOrderByUserProduct))
	e.GET("/order-buyer/:id", middleware.Auth(h.GetOrderByUserSeller))
	e.GET("/order-user", middleware.Auth(h.GetOrderbyUser))
	e.POST("/order", middleware.Auth(h.CreateOrder))
	e.DELETE("/order/:id", middleware.Auth(h.DeleteOrder))
	e.DELETE("/delete-order", middleware.Auth(h.DeleteAllOrder))

}
