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
  type SuccessResult struct {
  Status string `json:"status"`
  Data interface{} `json:"data"`
  }

  type ErrorResult struct {
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

- Create folder /models

  > File: `response.go`
  ```bash
  type Response struct {
	Wilayah   string              `json:"wilayah"`
	Perolehan []PerolehanResponse `json:"perolehan"`
  }

  type PerolehanResponse struct {
    Partai     string  `json:"partai"`
    TotalSuara float64 `json:"total_suara"`
  }
  ```

## Create data

- Create file data.go in /models

  > File `data.go`
  ```bash
  type PartaiData struct {
	Warna       string `json:"warna"`
	IsAceh      bool   `json:"is_aceh"`
	IdPilihan   int    `json:"id_pilihan"`
	NomorUrut   int    `json:"nomor_urut"`
	Nama        string `json:"nama"`
	NamaLengkap string `json:"nama_lengkap"`
  }
  ```

- Create folder /controllers and file data.go

  > file `data.go`

  ## func Read JSON

    ```bash
    func ReadJSON(filepath string, v interface{}) error {
    // Read file JSON
    file, err := os.Open(filepath)
    if err != nil {
      fmt.Println("Error read file:", err)
      return err
    }
    defer file.Close()

    // Parse JSON
    bytes, err := ioutil.ReadAll(file)
    if err != nil {
      fmt.Println("Error parsing JSON:", err)
      return err
    }

    err = json.Unmarshal(bytes, &v)
    if err != nil {
      fmt.Println("Error parsing JSON:", err)
      return err
    }
    return nil
    }
    ```

