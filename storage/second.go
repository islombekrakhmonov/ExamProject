package storage

import (
	"Proj/controller"
	"errors"
	"fmt"
)


var ReadShopCart =controller.ReaderShopCart()
var ReadProduct = controller.ReaderProduct()
var ReadUser = controller.ReaderUser()
var UserName string

func CheckUser() (string,error) {
	fmt.Println("Please enter your name")
	fmt.Scan(&UserName)
	for i := range ReadUser{
		if UserName == ReadUser[i].Name{
			return ReadUser[i].ID, nil
		} 
	} 
	return UserName, errors.New("you entered wrong name or your name is not in our system")
}

func ClientPurchaseHistory(){
	user, err := CheckUser()
    if err != nil {
		panic(err)
	} 
	for i,_ := range ReadShopCart{
		var priceOfProduct int 
		if user == ReadShopCart[i].UserID{
			for t:=0;t<len(ReadProduct);t++{
				if ReadShopCart[i].ProductID == ReadProduct[t].ID{
					priceOfProduct = ReadShopCart[i].Count * ReadProduct[t].Price
					fmt.Printf("Product name: %v; Price: %v; Count: %v; Total Price: %v; Time: %v\n",ReadProduct[t].Name, ReadProduct[t].Price, ReadShopCart[i].Count, priceOfProduct, ReadShopCart[i].Date)
					
				}
			}
		}
	}
}
