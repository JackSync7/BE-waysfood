package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	UserRoutes(e)
	AuthRoutes(e)
	TransactionRoutes(e)
	ProductRoutes(e)
	CartsRoutes(e)
	OrderRoutes(e)

}
