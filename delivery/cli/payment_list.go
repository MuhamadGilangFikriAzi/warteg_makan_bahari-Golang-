package cli

import (
	"WMB/model"
	"fmt"
)

func PaymentList(dataPayment []model.Transaction) {
	fmt.Printf("%-5s%-20s%-20s%-5s%10s\n", "id", "Nama Customer", "No Meja", "Total", "Status")
	for _, value := range dataPayment {
		fmt.Printf("%-5d%s%-20d%-5d%10s\n", value.Id, value.CustomerName, value.TableId, value.Total, value.Status)
	}
}
