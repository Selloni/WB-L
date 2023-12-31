// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package main

import "fmt"

func main() {
	var ss int64 = 5
	fmt.Println(setBit(ss, 1, 1)) // число в котором меняем бит/
}

func setBit(ss int64, i int, bit int) int64 {
	var tmp int64 = 1 // маска
	tmp = tmp << i    // сдвиг по маске
	if bit == 0 {     //  устанавливаемое значение
		return tmp ^ ss // если в маске единичка и в чиле единичка переводит его в ноль
	}
	return tmp | ss // унарынй или если в маске 1 возвращает возвращает 1
}
