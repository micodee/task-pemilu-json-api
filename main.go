package main

import (
	"fmt"
	"net/http"
	"partai/controllers"
	"partai/dto/result"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	// add route with method GET
	e.GET("/hasil", func(c echo.Context) error {
		// create map data with format JSON
		data := controllers.MapData()

		// send response with data JSON
		return c.JSON(http.StatusOK, result.SuccessResult{Status: "success", Data: data})
	})

	// create server port
	port := "8080"
	fmt.Println("server running on port", port)
	e.Logger.Fatal(e.Start("localhost:" + port))
}
