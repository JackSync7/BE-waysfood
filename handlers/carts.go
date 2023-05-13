package handlers

import (
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerCart struct {
	CartsRepository repositories.CartsRepository
}

func HandlerCarts(CartsRepository repositories.CartsRepository) *handlerCart {
	return &handlerCart{CartsRepository}
}

func (h *handlerCart) FindCarts(c echo.Context) error {
	carts, err := h.CartsRepository.FindCarts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: carts})
}

func (h *handlerCart) GetCarts(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	carts, err := h.CartsRepository.GetCarts(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: carts})
}

func (h *handlerCart) CreateCarts(c echo.Context) error {
	request := new(models.Carts)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// data form pattern submit to pattern entity db user
	carts := models.Carts{
		// Qty: request.Qty,
		// ProductID: request.ProductID,

	}

	data, err := h.CartsRepository.CreateCarts(carts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerCart) DeleteCarts(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	carts, err := h.CartsRepository.GetCarts(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CartsRepository.DeleteCarts(carts, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

//	func convertResponse(u models.User) usersdto.UserResponse {
//		return usersdto.UserResponse{
//			ID:        u.ID,
//			Fullname:  u.Fullname,
//			Email:     u.Email,
//			Password:  u.Password,
//			Subscribe: u.Subscribe,
//		}
//	}
// func DeleteResponse(u models.User) usersdto.UserDeleteResponse {
// 	return usersdto.UserDeleteResponse{
// 		ID: u.ID,
// 	}
// }
