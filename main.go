package main

import (
	"errors"
	"fmt"
	"math"
)

const IMPPower = 2

func main() {
	fmt.Println("-_- Калькулятор индекса массы тела -_-")
	for {
		userWeight, userHeight := getUserInput()
		IMT, err := calculateIMT(userWeight, userHeight)
		if err != nil {
			// fmt.Println("Не заданы параметры для рассчета")
			// continue
			panic("Не заданы параметры для рассчета")
		}
		outputResult(IMT)
		isReapeteCalculetion := checkRepeatCalculation()
		if !isReapeteCalculetion {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный диффицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас диффицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculateIMT(userWeight float64, userHeight float64) (float64, error) {
	if userWeight <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IMT := userWeight / math.Pow(userHeight/100, IMPPower)
	return IMT, nil
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

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать еще расчет (y/n): ")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
