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
	PaymentList(dataPayment)
	fmt.Print("Pilih id yang ingin dilakukan pembayaran : ")
	fmt.Scan(&pilihan)
	if paymentRepo.CheckUnpaymentTransaction(pilihan) {
		fmt.Print("Apakah Anda yakin ingin menyelesaikan pembayaran(y/n) : ")
		fmt.Scan(&pilihan2)
		if pilihan2 == "y" {
			paymentRepo.UpdatePayment(pilihan)
			fmt.Println("Pembayaran Berhasil...")
		} else {
			fmt.Println("pembayaran dibataklan...")
		}
	} else {
		fmt.Println("Tidak ada pembayaran dengan id tersebut...")
		utility.PauseProgram()
	}

}
