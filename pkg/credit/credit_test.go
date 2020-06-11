package credit_test // взяли пакет credit, добавили _test

import (
	"fmt"
	"go_hw2/pkg/credit"
)

func ExampleCalculate() { // имя функции - Example + имя проверяемой функции
	fmt.Println(credit.Calculate(1_000_000, 3, 20))
	fmt.Println(credit.Calculate(10_000, 3, 20))
	fmt.Println(credit.Calculate(3_000_000, 5, 20))
	// Output:
	// 3718400 33862400 133862400
	// 37184 338624 1338624
	// 7954830 177289800 477289800
}
