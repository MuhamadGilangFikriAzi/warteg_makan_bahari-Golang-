package delivery

import (
	"WMB/delivery/utility"
	"fmt"
)

func MainMenu() {
	utility.SetHeader("Warteg Makan Bahari")
	fmt.Println("1. Pesan Makanan")
	fmt.Println("2. List Makanan")
	fmt.Println("3. Cari Makanan ")
	fmt.Println("4. Pembayaran")
	fmt.Println("5. Keluar")
	fmt.Println("Pilih Menu 1-5")
}
