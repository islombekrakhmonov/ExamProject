package controller

import (
	"Proj/model"
	"encoding/json"
	"fmt"
	"os"
)

func ReaderShopCart() []model.ShopCart{
	data, err := os.ReadFile("../data/shop_cart.json")
	if err != nil{
		fmt.Println(err)
		return nil
	}
	var something []model.ShopCart
	err = json.Unmarshal(data, &something)
		if err != nil{
			fmt.Println(err)
		}
	return something
}