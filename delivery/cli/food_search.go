package cli

import (
	"WMB/delivery/utility"
	"WMB/manager"
	"fmt"
)

func SearchFood(usecase manager.UseCaseManager) {
	var search string
	utility.SetHeader("Cari Makanan")
	fmt.Print("Cari Makanan apa ? (Code / Nama makanan) : ")
	fmt.Scan(&search)
	food := usecase.FoodListUseCase().Search(search)
	if food.FoodName != "" {
		fmt.Printf("%s %s %s %d\n", "Ditemukan", food.FoodName, "Dengan Harga", food.Price)
	}
}
