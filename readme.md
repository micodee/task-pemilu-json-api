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

  - func Read JSON

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

  - func MapData

    ```bash
    func MapData() []models.Response {
    res := []models.Response{}

    var listWilayah map[string]interface{}
    err := ReadJSON("data/0.json", &listWilayah)
    if err != nil {
      return []models.Response{}
    }

    var partai map[string]models.PartaiData
    err = ReadJSON("data/partai.json", &partai)
    if err != nil {
      return []models.Response{}
    }	

    var listDPR map[string]interface{}
    err = ReadJSON("data/dprri.json", &listDPR)
    if err != nil {
      return []models.Response{}
    }

    for iWilayah, v := range listWilayah {
      data := models.Response{}
      
      for i, x := range v.(map[string]interface{}) {
        if i == "nama" {
          data.Wilayah = x.(string)
        }
      }

      for iDPR, vDPR := range listDPR {

        if iDPR == "table" {
          for iTable, vTable := range vDPR.(map[string]interface{}) {
            if iTable == iWilayah {
              for iTBValue, vTBValue := range vTable.(map[string]interface{}) {
                for _, pt := range partai {
                  idPartai, _ := strconv.Atoi(iTBValue)
                  if pt.IdPilihan == idPartai {
                    perolehan := models.PerolehanResponse{}

                    perolehan.Partai = pt.Nama
                    perolehan.TotalSuara = vTBValue.(float64)

                    data.Perolehan = append(data.Perolehan, perolehan)
                  }
                }
              }
            }
          }
        }

      }

      res = append(res, data)
    }

    return res
    }
    ```

  - Update data in main.go & add CORS Origin

    > file `main.go`

      ```bash
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
      ```
