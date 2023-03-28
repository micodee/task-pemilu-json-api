package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"partai/models"
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

func Dapil() {
	var dapilMap map[string]models.DapilData
	err := ReadJSON("data/0.json", &dapilMap)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// melakukan looping pada setiap nilai dalam map
	for _, value := range dapilMap {
		data := models.Response{}
		data.Wilayah = value.Nama
		fmt.Println("Wilayah:", value.Nama)
		// melakukan looping pada nilai dalam setiap elemen struct DapilData
		for _, v := range value.Dapil {
			fmt.Println("Dapil:", v)
		}
	}
}

func Partai() {
	var partaiMap map[string]models.PartaiData
	err := ReadJSON("data/partai.json", &partaiMap)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// melakukan looping pada setiap elemen map
	for key, value := range partaiMap {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}

func DprRI() {
	var dprMap map[int]models.DprData
	err := ReadJSON("data/dprri.json", &dprMap)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// melakukan looping pada setiap elemen map
	for key, value := range dprMap {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
