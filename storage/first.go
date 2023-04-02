package storage

import (
	"Proj/model"
	"fmt"
)

func FilterDates(from, to string){
	newDate :=[]model.ShopCart{}

	for i:= 0; i<len(ReadShopCart); i++{
		if 	ReadShopCart[i].Date >= from && ReadShopCart[i].Date <= to{
			newDate = append(newDate, ReadShopCart[i])
	    }
	}
	
	for i:=0;i<len(newDate);i++{
		fmt.Printf("ID: %v; ProductId: %v; UserId: %v; Count: %v; Date: %v\n",newDate[i].ID,newDate[i].ProductID,newDate[i].UserID,newDate[i].Count,newDate[i].Date)
	}
}
