package storage

import (
	"fmt"
)

func ClientTolalSpending(){
	user, err := CheckUser()
    if err != nil {
		panic(err)
	} 
	var priceOfProduct int
	var totalPrice int 
	for i,_ := range ReadShopCart{
		if user == ReadShopCart[i].UserID{
			for t:=0;t<len(ReadProduct);t++{
				if ReadShopCart[i].ProductID == ReadProduct[t].ID{
					priceOfProduct = ReadShopCart[i].Count * ReadProduct[t].Price
					totalPrice += priceOfProduct
					
				}
			}
		}
	}
	fmt.Printf("Name: %v Total Spenditure: %v \n",UserName,totalPrice)
}