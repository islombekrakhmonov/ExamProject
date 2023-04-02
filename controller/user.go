package controller

import (
	"Proj/model"
	"encoding/json"
	"fmt"
	"os"
)

func ReaderUser() []model.User{
	data, err := os.ReadFile("../data/user.json")
	if err != nil{
		fmt.Println(err)
		return nil
	}
	var something []model.User
	err = json.Unmarshal(data, &something)
		if err != nil{
			fmt.Println(err)
		}
	return something
}