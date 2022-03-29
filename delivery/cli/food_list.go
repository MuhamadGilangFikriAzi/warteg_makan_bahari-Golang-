package cli

import (
	"WMB/delivery/utility"
	"WMB/usecase"
	"fmt"
)

func FoodList(food usecase.FoodListUseCase) {
	utility.SetHeader("List Makanan")
	formatHeaderTable := "%-5s%-20s%-20s\n"
	formatTable := "%-5d%-20s%-20d\n"
	fmt.Printf(formatHeaderTable, "No", "Nama Makanan", "Harga")
	for idx, data := range food.AllFood() {
		fmt.Printf(formatTable, idx+1, data.FoodName, data.Price)
	}
}
