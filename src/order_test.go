package src

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/require"
)

func MockOrderData() {
	menuData := []MenuData{
		{Id: 1, FoodName: "Subway Club Salad", TotalOrder: 0},
		{Id: 2, FoodName: "Potato Chip", TotalOrder: 0},
		{Id: 3, FoodName: "Chocolate Chunk Cookie", TotalOrder: 0},
		{Id: 4, FoodName: "Cool Wrap", TotalOrder: 0},
		{Id: 5, FoodName: "Cobb Salad", TotalOrder: 0},
	}

	for _, value := range menuData {
		MenuStore[value.Id] = value
	}
}

func MockUserData() {
	menuData := []UserData{
		{Id: 1, UserName: "Alfie"},
		{Id: 2, UserName: "Freddy"},
		{Id: 3, UserName: "George"},
		{Id: 4, UserName: "Henrie"},
		{Id: 5, UserName: "John"},
	}

	for _, value := range menuData {
		UserStore[value.Id] = value
	}
}

func TestSetUp(t *testing.T) {
	MockOrderData()
	MockUserData()
}

func TestPlaceOrder(t *testing.T) {
	args := []Order{
		{FoodId: "4", UserId: "1"},
		{FoodId: "4", UserId: "2"},
		{FoodId: "2", UserId: "2"},
		{FoodId: "3", UserId: "1"},
		{FoodId: "4", UserId: "5"},
		{FoodId: "1", UserId: "5"},
		{FoodId: "2", UserId: "3"},
		{FoodId: "5", UserId: "5"},
		{FoodId: "5", UserId: "3"},
		{FoodId: "4", UserId: "4"},
		{FoodId: "2", UserId: "1"},
		{FoodId: "4", UserId: "3"},
		{FoodId: "2", UserId: "5"},
	}

	result, err := PlaceOrder(args)

	fmt.Println("err", err != nil)
	if err != nil {
		require.EqualError(t, err, "duplicate order placed", "unexpected error message")
	} else {
		require.Equal(t, result[0].Id, MenuStore[4].Id, "Top One is Mismatched")
		require.Equal(t, result[1].Id, MenuStore[2].Id, "Top One is Mismatched")
		require.Equal(t, result[2].Id, MenuStore[5].Id, "Top One is Mismatched")
	}
}

func TestDuplicatePlaceOrder(t *testing.T) {
	args := []Order{
		{FoodId: "4", UserId: "1"},
		{FoodId: "3", UserId: "3"},
		{FoodId: "2", UserId: "2"},
	}

	_, err := PlaceOrder(args)

	if err != nil {
		require.EqualError(t, err, "duplicate order placed", "unexpected error message")
	}
}
