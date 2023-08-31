package handler

import (
	"fakhry/cleanarch/app/middlewares"
	"fakhry/cleanarch/features/item"
	"fakhry/cleanarch/helpers"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	itemService item.ItemServiceInterface
}

func New(service item.ItemServiceInterface) *ItemHandler {
	return &ItemHandler{
		itemService: service,
	}
}

func (handler *ItemHandler) GetAllItem(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	fmt.Println("id:", idToken)
	if idToken != 1 {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, "unauthorized", nil))
	}

	nameQuery := c.QueryParam("name")
	result, err := handler.itemService.GetAll(nameQuery)
	if err != nil {
		// return c.JSON(http.StatusInternalServerError, map[string]any{
		// 	"code": 500,
		// 	"message": "hello world",
		// })
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	// mapping dari struct core to struct response
	// var usersResponse []UserResponse
	// for _, value := range result {
	// 	usersResponse = append(usersResponse, UserResponse{
	// 		ID:        value.ID,
	// 		Name:      value.Name,
	// 		Email:     value.Email,
	// 		Address:   value.Address,
	// 		CreatedAt: value.CreatedAt,
	// 	})
	// }
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", result))
	// return c.JSON(http.StatusOK, map[string]any{
	// 	"code":    200,
	// 	"message": "success read data",
	// 	"data":    usersResponse,
	// })

}
