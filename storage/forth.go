package storage

import (
	"fmt"
)

func GetTheNumberOfProductSold(){
	countr := make(map[string]int)

	for i,_ := range ReadShopCart{
		for t:=0;t<len(ReadProduct);t++{
		    if ReadShopCart[i].ProductID == ReadProduct[t].ID{
			   countr[ReadProduct[t].Name]= ReadShopCart[i].Count
		    }
	    }	
	}
	for p,l := range countr{
		fmt.Println("Name:" ,p, "count:", l)
	}
}