package router

import (
	"net/http"
	"sekolahApp/app/middlewares"
	_userData "sekolahApp/features/users/data"
	_userHandler "sekolahApp/features/users/handler"
	_userService "sekolahApp/features/users/service"
	"sekolahApp/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c *echo.Echo) {
	UserData := _userData.NewUsersQuery(db)
	UserService := _userService.NewUsersLogic(UserData)
	UserHandlerAPI := _userHandler.NewUsersHandler(UserService)

	c.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "get test success", nil))
	})

	c.POST("/login", UserHandlerAPI.Login)
	c.GET("/users", UserHandlerAPI.GetAllUser, middlewares.JWTMiddleware())
	c.POST("/users", UserHandlerAPI.CreateUser, middlewares.JWTMiddleware())
	c.DELETE("/user/:user_id", UserHandlerAPI.DeleteUser, middlewares.JWTMiddleware())
	c.GET("/profile", UserHandlerAPI.GetUser, middlewares.JWTMiddleware())
	c.PUT("/update", UserHandlerAPI.UpdateUser, middlewares.JWTMiddleware())
}
