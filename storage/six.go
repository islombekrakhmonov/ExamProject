package storage

import (
	"fmt"
	"sort"
)

func GetTheLeastSelled(){
	var numberOfS []int
	countr := make(map[string]int)

	for i,_ := range ReadShopCart{
		for t:=0;t<len(ReadProduct);t++{
		    if ReadShopCart[i].ProductID == ReadProduct[t].ID{
			   countr[ReadProduct[t].Name]= ReadShopCart[i].Count
		    }
	    }	
	}
	for _,l := range countr{
		numberOfS = append(numberOfS, l)
	} 

	smallest, biggest := numberOfS[0], numberOfS[0]

	for _, v := range numberOfS {
		if v > biggest {
			biggest = v
		}
		if v < smallest {
			smallest = v
		}
	}
	sort.Slice(numberOfS, func(i, j int) bool {
		return numberOfS[i] < numberOfS[j]
	})

	for i:=0; i<10; i++{
		for t,v := range countr{
			if numberOfS[i] == v{
				fmt.Println(t, numberOfS[i])
			}
		}
	}
}