package delivery

import (
	"WMB/delivery/utility"
	"WMB/manager"
	"fmt"
)

func Payment(usecase manager.UseCaseManager) {
	var pilihan int
	var pilihan2 string
	utility.SetHeader("Pembayaran")
	paymentRepo := usecase.PaymentUseCase()
	dataPayment := paymentRepo.AllPayment()
	fmt.Printf("%-5s%-20s%-20s%-5s%10s\n", "id", "Nama Customer", "No Meja", "Total", "Status")
	for _, value := range dataPayment {
		fmt.Printf("%-5d%-20s%-20d%-5d%10s\n", value.Id, value.CustomerName, value.TableId, value.Total, value.Status)
	}
	fmt.Print("Pilih id yang ingin dilakukan pembayaran : ")
	fmt.Scan(&pilihan)
	fmt.Print("Apakah Anda yakin ingin menyelesaikan pembayaran(y/n) : ")
	fmt.Scan(&pilihan2)
	if pilihan2 == "y" {
		usecase.PaymentUseCase().UpdatePayment(pilihan)
		fmt.Println("Pembayaran Berhasil...")
	} else {
		fmt.Println("pembayaran dibataklan...")
	}

}
