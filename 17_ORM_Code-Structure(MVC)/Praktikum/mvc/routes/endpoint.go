package routes

import (
	"github.com/labstack/echo/v4"
	"mvc/controllers"
)

func New() *echo.Echo{
	e := echo.New()
	//User
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	//Buku
	e.GET("/books", controllers.GetBookController)
	e.GET("/books/:id", controllers.GetBookidController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	return e
}