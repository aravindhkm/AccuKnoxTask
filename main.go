package main

import (
	"fmt"
	"github.com/aravindhkm/AccuKnoxTask/src"
)

func main() {
	src.ConfigOrderData()
	src.ConfigUserData()
	args := src.GetLogFileData()
	result, err := src.PlaceOrder(args)

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {		
		for index, menuValue := range result {
			fmt.Printf("Top (%d) Id: %v, FoodName: %v, TotalOrder: %v  \n", index+1, menuValue.Id, menuValue.FoodName, menuValue.TotalOrder)
		}
	} 
}
