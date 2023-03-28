package main

import (
	"fmt"
	"net/http"
	"partai/controllers"
	"partai/dto/result"
	"partai/models"

	"github.com/labstack/echo/v4"
)

func main() {
	controllers.Dapil()
	controllers.Partai()
	controllers.DprRI()

	e := echo.New()

	// add route with method GET
	e.GET("/hasil", func(c echo.Context) error {
		// create map data with format JSON
		data := models.Response{
			Wilayah: models.DapilData.Nama,
		}

		// send response with data JSON
		return c.JSON(http.StatusOK, result.SuccessResult{Status: "get success", Data: data})
	})

	// create server port
	port := "8080"
	fmt.Println("server running on port", port)
	e.Logger.Fatal(e.Start("localhost:" + port))
}
