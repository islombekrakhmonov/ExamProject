package controller

import (
	"Proj/model"
	"encoding/json"
	"fmt"
	"os"
)

func ReaderProduct() []model.Product {
	data, err := os.ReadFile("../data/product.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var something []model.Product
	err = json.Unmarshal(data, &something)
	if err != nil {
		fmt.Println(err)
	}
	return something
}
