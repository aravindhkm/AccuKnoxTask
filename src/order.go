package src

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	_ "reflect"
	"sort"
	"strconv"
	"strings"
)

type UserData struct {
	Id       int
	UserName string
}

type MenuData struct {
	Id         int
	FoodName   string
	TotalOrder int
}

type Order struct {
	FoodId string
	UserId string
}

var MenuStore = make(map[int]MenuData)
var UserStore = make(map[int]UserData)

func ConfigOrderData() {
	var menuData []MenuData // 15 items
	// menu file read
	orderBytes, err := os.ReadFile("order.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(orderBytes, &menuData)
	if err != nil {
		panic(err)
	}

	for _, value := range menuData {
		MenuStore[value.Id] = value
	}
}

func ConfigUserData() {
	var userData []UserData // 5 users
	// user file read
	userBytes, err := os.ReadFile("user.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(userBytes, &userData)
	if err != nil {
		panic(err)
	}

	for _, value := range userData {
		UserStore[value.Id] = value
	}
}

func PlaceOrder(args []Order) ([]MenuData, error) {
	var orderCounts = make(map[string]int)

	for _, value := range args {
		key := value.UserId + "||" + value.FoodId

		if orderCounts[key] == 0 {
			foodIdInt, _ := strconv.Atoi(value.FoodId)

			if thisMenu, ok := MenuStore[foodIdInt]; ok {
				thisMenu.TotalOrder++
				MenuStore[foodIdInt] = thisMenu
			}

		} else {
			return nil, errors.New("duplicate order placed")
		}
		orderCounts[key]++
	}

	menuSort := make([]MenuData, 0, len(MenuStore))

	for _, structValue := range MenuStore {
		menuSort = append(menuSort, structValue)
	}

	sort.Slice(menuSort, func(i, j int) bool {
		return menuSort[i].TotalOrder > menuSort[j].TotalOrder
	})

	var topThree []MenuData

	if len(menuSort) > 3 {
		topThree = menuSort[0:3]
	} else {
		topThree = menuSort
	}

	return topThree, nil
}

func GetLogFileData() []Order {
	var args []Order

	logFileBytes, err := os.Open("log")
	if err != nil {
		panic(err)
	}
	defer logFileBytes.Close()

	scanner := bufio.NewScanner(logFileBytes)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		userId := parts[0]
		foodId := parts[1]

		args = append(args, Order{UserId: userId, FoodId: foodId})
	}

	return args
}
