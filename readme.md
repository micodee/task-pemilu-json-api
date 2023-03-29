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

- Save (ctrl+s) for import automatis

  ```bash
  go mod tidy
  ```