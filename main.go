package main

import (
	"fmt"
	"math"
)

func main() {
	const IMPPower = 2
	var userHeight, userWeight float64
	fmt.Print("-_- Калькулятор индекса массы тела -_-\n")
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&userWeight)
	var IMT = userWeight / math.Pow(userHeight, IMPPower)
	fmt.Print("Ваш индекс массы тела: ")
	fmt.Print(IMT)
}
