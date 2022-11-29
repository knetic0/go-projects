package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"net/http"
)

type User struct {
	username string `json:'username'`
	name     string `json:'name'`
	surname  string `json:'surname'`
}

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "mainHandler fonksiyonu basarili bir sekilde calisti!.")
}

func userHandler(c echo.Context) error {
	dataType := c.Param("data")

	username := c.QueryParam("username")
	name := c.QueryParam("name")
	surname := c.QueryParam("surname")

	if dataType == "string" {

		return c.String(http.StatusOK, fmt.Sprintf("Username : %s Name : %s Surname : %s", username, name, surname))

	} else if dataType == "json" {

		return c.JSON(http.StatusOK, map[string]string{
			"username": username,
			"name":     name,
			"surname":  surname,
		})
	}

	return c.String(http.StatusBadRequest, "Yalnizca JSON veya String Parametreleri kabul edilebilir.")

}

func addUser(c echo.Context) error {

	user := User{}

	body, error := ioutil.ReadAll(c.Request().Body)

	if error != nil {
		return error
	}

	error = json.Unmarshal(body, &user)

	if error != nil {
		return error
	}

	fmt.Println(user)

	fmt.Println(user.username)
	fmt.Println(user.name)
	fmt.Println(user.surname)
	return c.String(200, "Basarili!")
}

func userMessage(c echo.Context) error {

	return c.String(200, "Welcome to System. Hello World!")
}

func sendUser(c echo.Context) error {

	newUser := User{"knetic", "Mehmet", "SOLAK"}
	return c.String(200, fmt.Sprintf("%s", newUser))

}

func takeUser(c echo.Context) error {

	bytes_json, bytes_err := ioutil.ReadAll(c.Request().Body)

	if bytes_err != nil {
		return bytes_err
	}

	var user User
	json_error := json.Unmarshal(bytes_json, &user)

	if json_error != nil {
		return json_error
	}

	fmt.Println(user)
	return c.String(200, "Basarili")

}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Admin endpointindeyiz.")
}

func main() {

	fmt.Println("Hello World")

	e := echo.New()
	// e.Use(middleware.Logger()) // bu komut ile butun endpointlere logger ekleyebiliriz.
	// e.use(middleWare.LoggerWithConfig(middleware.LoggerConfig{Format: "method=${method}, url=${url}, status=${status}",}))
	// bu sekilde ozellestirilmis bir sekilde log'lari yazdirabiliriz.

	e.GET("/main", mainHandler, middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method = ${method}, url = ${url}, status = ${status}\n",
	})) //endpointler api uzerinde belirli bir amac icin olusturulmus metoda verilen isimdir.

	e.GET("/user/:data", userHandler)
	e.GET("/user", userMessage)

	adminGroup := e.Group("/admin")

	adminGroup.Use(middleware.Logger())

	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "123" {
			return true, nil
		}

		return false, nil
	}))

	adminGroup.GET("/main", mainAdmin)

	/*
		e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			// Be careful to use constant time comparison to prevent timing attacks
			if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
				return true, nil
			}
			return false, nil
		}))
	*/

	e.POST("/user", addUser)
	e.GET("/users", sendUser)
	e.POST("/users", takeUser)

	e.Start(":8080")
}
