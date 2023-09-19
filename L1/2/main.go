package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	ExampleOne()
	fmt.Println("-----")
	ExampleTwo()
}

// Сделано с помощью мьютекса,
// обычный счетчик, который блокирует мьютекс
func ExampleOne() {
	var wg sync.WaitGroup
	nums := []int{2, 4, 6, 8, 10}
	for _, i := range nums {
		wg.Add(1)
		go func(i int) {
			res := math.Pow(float64(i), 2)
			fmt.Println(res)
			defer wg.Done()
		}(i)
		wg.Wait()
	}
} //
// - не встраивайте мьютекс в структуру
// - не храните ссылку на мьютекс в поле структуры
// - методы с мьютекс, должны иметь ссылочный ресивер
/////////////////////////////////////////////////////

// С помощью каналов
func ExampleTwo() {
	nums := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	go func() {
		for _, i := range nums {
			ch <- i // запись в канал
		}
		close(ch)
	}() // этот блок кода в горутине, выполняеться в дргуом ядре
	for cc := range ch { // считываю с канала, пока он не будет закрыт
		res := math.Pow(float64(cc), 2)
		fmt.Println(res)
	}
} //
//Запись в неициализированный канал блокирует поток навсегда;
//Чтение из неинициализированного канала блокирует поток навсегда;
//Запись в закрытый канал вызывает панику;
//Чтение из закрытого канала даёт нулевое значение мгновенно.