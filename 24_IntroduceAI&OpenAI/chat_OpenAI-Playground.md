USER
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

ASSISTANT
This code is a function that creates a new user in a database. 

1. It takes an echo.Context object as a parameter.
2. It creates an empty User object from the model.
3. It binds the data from the request to the User object.
4. It attempts to create the user in the database using the DB object from the configs package.
5. If there is an error, it returns a HTTP error with the error message.
6. If the user is successfully created, it returns a JSON response with a success message and the user object.

USER
make a different code
ASSISTANT
func UpdateUserController(c echo.Context) error {
    userID := c.Param("id")
    user := model.User{}
    if err := configs.DB.First(&user, userID).Error; err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "User not found")
    }
    c.Bind(&user)
    if err := configs.DB.Save(&user).Error; err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "success update user",
        "user":    user,
    })
}

This code is a function that updates an existing user in a database.

1. It retrieves the user ID from the request parameters.
2. It creates an empty User object from the model.
3. It attempts to find the user with the given ID in the database using the DB object from the configs package.
4. If the user is not found, it returns a HTTP error with a "User not found" message.
5. It binds the updated data from the request to the User object.
6. It attempts to save the updated user in the database using the DB object from the configs package.
7. If there is an error, it returns a HTTP error with the error message.
8. If the user is successfully updated, it returns a JSON response with a success message and the updated user object.

