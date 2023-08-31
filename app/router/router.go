package router

import (
	"fakhry/cleanarch/app/middlewares"
	_userData "fakhry/cleanarch/features/user/data"
	_userHandler "fakhry/cleanarch/features/user/handler"
	_userService "fakhry/cleanarch/features/user/service"

	_itemData "fakhry/cleanarch/features/item/data"
	_itemHandler "fakhry/cleanarch/features/item/handler"
	_itemService "fakhry/cleanarch/features/item/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	itemData := _itemData.New(db)
	itemService := _itemService.New(itemData)
	itemHandlerAPI := _itemHandler.New(itemService)

	e.POST("/login", userHandlerAPI.Login)
	e.GET("/users", userHandlerAPI.GetAllUser, middlewares.JWTMiddleware())
	e.GET("/users/:user_id", userHandlerAPI.GetUserById)
	e.POST("/users", userHandlerAPI.CreateUser)

	e.GET("/items", itemHandlerAPI.GetAllItem, middlewares.JWTMiddleware())

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "hello world",
		})
	})
}
