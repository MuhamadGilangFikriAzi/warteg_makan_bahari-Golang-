package utility

import (
	"fmt"
	"os"
	"strings"
	"time"
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
		SayGoodbay()
	}
}

func SayGoodbay() {
	fmt.Println("Selamat tinggal, Terimakasih...")
	os.Exit(0)
}

func PauseProgram() {
	duration := time.Duration(10)
	time.Sleep(duration)
}

func GetTimestamp() string {
	return time.Now().String()
}
