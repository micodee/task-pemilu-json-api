package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"partai/models"
	"strconv"
)

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
									data.Perolehan.Partai = pt.Nama
									data.Perolehan.TotalSuara = vTBValue.(float64)
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