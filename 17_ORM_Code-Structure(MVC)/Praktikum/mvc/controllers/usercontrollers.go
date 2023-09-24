package controllers

import (
	"net/http"
	"mvc/configs"
	"mvc/model"
	"strconv"
	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
    var users []model.User
    if err := configs.DB.Find(&users).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success get all users",
        "users":   users,
    })
}

func GetUserController(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
    }
    var user []model.User
    if err := configs.DB.First(&user, id).Error; err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "User not found")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success get user",
        "user":    user,
    })
}

func CreateUserController(c echo.Context) error {
    user := model.User{}
    c.Bind(&user)
    if err := configs.DB.Create(&user).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success create new user",
        "user":    user,
    })
}

func DeleteUserController(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
    }
	var user model.User
    if err := configs.DB.Delete(&user, id).Error; err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
    }
    return c.JSON(http.StatusOK, map[string]string{
        "message": "User deleted",
    })
}

func UpdateUserController(c echo.Context) error {
    userid, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
    }
    var user [] model.User
    if err := configs.DB.First(&user, userid).Error; err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "User not found")
    }
    c.Bind(&user)
    if err := configs.DB.Save(&user).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "User updated",
        "user":    user,
    })
}