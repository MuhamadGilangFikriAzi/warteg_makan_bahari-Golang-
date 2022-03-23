package utility

import (
	"fmt"
	"os"
	"strings"
)

func SetHeader(title string) {
	fmt.Println(title)
	fmt.Println(strings.Repeat("*", 50))
}

func ChoiceToMainMenu() {
	var chose string
	fmt.Println("kembali ke main menu (y/n)")
	fmt.Scan(&chose)
	if chose != "y" {
		fmt.Println("Selamat tinggal, Terimakasih...")
		os.Exit(0)
	}
}
