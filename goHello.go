package main

import (
	"fmt"
	"strconv"
)

func main() {
	lastWeekTemp := [...]int{2: 4, 6: 1}

	sumTemp := 0

	for i := 0; i < len(lastWeekTemp); i++ {
		sumTemp += lastWeekTemp[i]
		fmt.Println("summ: " + strconv.Itoa(sumTemp))
	}

	avarage := float32(sumTemp / len(lastWeekTemp))

	fmt.Println(sumTemp)
	fmt.Println(len(lastWeekTemp))

	fmt.Println(avarage)

	m := make(map[string]string) // создаём map — о применении функции make для создания переменных типа map будет рассказано ниже
	m["foo"] = "bar"             // заполняем парами «ключ-значение»
	m["ping"] = "pong"
	fmt.Println(m) // печатаем
}
