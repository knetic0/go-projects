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
	value string
}

func MainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func AdminMainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Admin. Welcome to the System!")
}

func MainHandlerScrap(c echo.Context) error {
	url, err := http.Get("http://localhost:8080/main")
	if err != nil {
		return err
	}
	defer url.Body.Close()

	response, err := ioutil.ReadAll(url.Body)
	if err != nil {
		return err
	}

	var str string
	json.Unmarshal(response, &str)

	fmt.Println("Suan da mainhandlerscrap icerisindeyim.")
	fmt.Println(str)

	return c.String(200, "Basarili")
}

func main() {
	fmt.Println("Hello World!")

	e := echo.New()
	e.GET("/main", MainHandler)
	e.GET("/mainscraping", MainHandlerScrap)

	admin_group := e.Group("/admin", middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time}, method=${method}, status=${status}\n",
	}))

	admin_group.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "mehmet" && password == "123" {
			return true, nil
		}
		return false, nil
	}))

	admin_group.GET("/main", AdminMainHandler)

	e.Start(":8080")
}
