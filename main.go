package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("-_- Калькулятор индекса массы тела -_-")
	userWeight, userHeight := getUserInput()
	IMT := calculateIMT(userWeight, userHeight)
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)
}

func calculateIMT(userWeight float64, userHeight float64) float64 {
	const IMPPower = 2
	IMT := userWeight / math.Pow(userHeight/100, IMPPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&userWeight)
	return userWeight, userHeight
}
