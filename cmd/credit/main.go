package main

import (
	"fmt"
	"go_hw2/pkg/credit"
)

func main() {
	payment, overpayment, amount := credit.Calculate(1_000_000, 3, 20)
	fmt.Println(payment, overpayment, amount)
}
