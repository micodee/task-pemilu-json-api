# Prepare

## Create main.go

- Init Golang

  ```bash
  go mod init partai
  ```

- Create file main.go

  > File: `main.go`

  - func main() {

    ```bash
    e := echo.New()

    // create server port
    port := "8080"
    fmt.Println("server running on port", port)
    e.Logger.Fatal(e.Start("localhost:" + port))
    ```

    }

- Save (ctrl+s) for import automatis and then

  ```bash
  go mod tidy
  ```

## Create response

- Create Folder /dto/result

  > File: `result.go`

  ```bash
  - type SuccessResult struct {
  Status string `json:"status"`
  Data interface{} `json:"data"`
  }

  - type ErrorResult struct {
  Status int `json:"status"`
  Message string `json:"message"`
  }
  ```

- Add route before port
  > File: `main.go`

  ```bash
  // add route with method GET
  e.GET("/hasil", func(c echo.Context) error {
  // create map data with format JSON
  data := "wilayah"

  // send response with data JSON
  return c.JSON(http.StatusOK, result.SuccessResult{Status: "success", Data: data})
  })
  ```

- Testing method GET in localhost:8080/hasil
  ```bash
  go run .
  ```
