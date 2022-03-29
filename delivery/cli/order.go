package cli

import (
	"WMB/delivery/utility"
	"WMB/manager"
	"WMB/model"
	"fmt"
)

func MakeOrder(usecase manager.UseCaseManager) {
	var customerName string
	var chouseTable int
	var status string
	var totalOrder int
	var chouseFood int
	var qty int
	var detail []model.TransactionDetail
	var total int

	avaible := usecase.AvaibleTableUseCase().AllAvaibleTable()
	utility.SetHeader("List Meja kosong")
	fmt.Printf("%-8s%-20s\n", "No Meja", "Status")
	for _, table := range avaible {
		if table.IsOrdered == false {
			status = "Kosong"
		} else {
			status = "Dipesan"
		}
		fmt.Printf("%-5d%-20s\n", table.TableNumber, status)
	}
	fmt.Print("Pilih Meja : ")
	fmt.Scan(&chouseTable)
	isBooked := false
	for _, tableBooked := range usecase.AvaibleTableUseCase().BookedTable() {
		if tableBooked.Id == chouseTable {
			isBooked = true
		}
	}
	if !isBooked {
		fmt.Print("Masukan Nama Customer : ")
		fmt.Scan(&customerName)
		FoodList(usecase.FoodListUseCase())
		fmt.Println("Berapa Pesanan? : ")
		fmt.Scan(&totalOrder)
		for i := 0; i < totalOrder; i++ {
			fmt.Print("Makanan ke-", i+1, "(no) : ")
			fmt.Scan(&chouseFood)
			fmt.Print("Jumlah makanan ke-", i+1, " : ")
			fmt.Scan(&qty)
			detail = append(detail, model.TransactionDetail{
				FoodId: chouseFood,
				Qty:    qty,
			})
		}

		for _, food := range detail {
			total += usecase.FoodListUseCase().GetPrice(food.FoodId) * food.Qty
		}
		dataTransaction := model.Transaction{
			Status:       "Proses",
			TableId:      chouseTable,
			CustomerName: customerName,
			Total:        total,
		}
		var decision string
		fmt.Print("Apakah Anda yakin memesan makanan ini ?(y/n) : ")
		fmt.Scan(&decision)
		if decision == "y" {
			transactionId := usecase.FoodOrderUseCase().InsertTransaction(dataTransaction)
			usecase.FoodOrderDetailUseCase().InsertTransactionDetail(transactionId, detail)
			usecase.AvaibleTableUseCase().OrderTable(chouseTable)

			fmt.Println("Terima kasih, makanan akan segera diantar kemeja anda...")
		}
	}
	fmt.Println("Meja yang anda pesan tidak tersedia...")
}
