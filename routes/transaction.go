package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	orderRepository := repositories.RepositoryOrder(mysql.DB)
	cartRepository := repositories.RepositoryCarts(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository, orderRepository, cartRepository)

	e.GET("/transactions", h.FindTransaction)
	e.GET("/transaction/:id", h.GetTransaction)
	e.GET("/transaction-user", middleware.Auth(h.GetUserTransaction))
	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	e.DELETE("/transaction/:id", middleware.Auth(h.DeleteTransaction))
	e.POST("/notification", h.Notification)

}
