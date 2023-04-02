package main

import (
	"Proj/storage"
	"fmt"
)

func main() {
	var userInput1 = GetUserInput1()
	switch userInput1{
	case "1": storage.FilterDates("2022-02-19","2022-02-26")
	case "2": storage.ClientPurchaseHistory()
	case "3": storage.ClientTolalSpending()
	case "4": storage.GetTheNumberOfProductSold()
	case "5": storage.GetTheMostSelled()
	case "6": storage.GetTheLeastSelled()
	}
}

func GetUserInput1()string{
	var userText string
	fmt.Println("Welcome")
	fmt.Println("Please choose what to do")
	fmt.Println("1. Filter by dates")
	fmt.Println("2. Show client purchase history")
	fmt.Println("3. Show client total spending")
	fmt.Println("4. Show number of product sold ")
	fmt.Println("5. Show the top selling products ")
	fmt.Println("6. Show the least selling products ")
	// fmt.Println("7. Show the top selling products ")
	// fmt.Println("8. Show the products sold by category ")
	// fmt.Println("9. Show the client who spent most ")
	// fmt.Println("10. Discount if client's buying more than 9 products ")
	fmt.Scan(&userText)
	return userText
}