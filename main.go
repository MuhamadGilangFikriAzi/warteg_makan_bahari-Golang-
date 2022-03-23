package main

import (
	"WMB/config"
	"WMB/delivery"
	"WMB/delivery/utility"
	"fmt"
	"os"
)

func main() {
	appConfig := config.NewConfig()
	for {
		var choice int
		delivery.MainMenu()
		fmt.Scan(&choice)
		switch choice {
		case 1:
			delivery.MakeOrder(appConfig.UseCaseManager)
		case 2:
			delivery.FoodList(appConfig.UseCaseManager.FoodListUseCase())
			utility.ChoiceToMainMenu()
		case 3:
			delivery.SearchFood(appConfig.UseCaseManager)
		case 4:
			delivery.Payment(appConfig.UseCaseManager)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Tidak ada pilihan itu")
		}
	}
}
