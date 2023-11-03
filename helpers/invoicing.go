package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateInvoiceNumber() string {
	timestamp := time.Now().Unix()
	random_num := rand.Intn(5000)
	// this is for unit testing (because if random, the invoice generated when creating real and testing has differenct random num)
	// random_num := 1
	invoiceNumber := fmt.Sprintf("%d-%d", timestamp, random_num)
	return invoiceNumber
}
