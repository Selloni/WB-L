package main

//Реализовать все возможные способы остановки выполнения горутины.

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// используем канал для остановки горутины
func exOne() {
	// Создание канала для передачи сигналов об остановке горутины
	stop := make(chan bool)
	// Запуск горутины

	go func() {
		for {

			select {
			default:
				// Выполнение работы в горутине
				fmt.Println("Канал...")
				time.Sleep(1 * time.Second)
			case <-stop:
				// Получен сигнал об остановке
				fmt.Println("Останавливаем через канал")
				return
			}
		}
	}()
	// Ждем 3 секунд, затем отправляем сигнал об остановке
	time.Sleep(3 * time.Second)
	stop <- false

	// Ждем завершения горутины
	time.Sleep(1 * time.Second)
	fmt.Println("Программа завершена через канал")
}

// используем контекст с функцией отмены.
func exTwo() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			// поулчаем сигнал для остановки
			case <-ctx.Done():
				fmt.Println("Останавливаем через контекст")
				return
			default:
				fmt.Println("Контекст...")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	cancel() // освобождаем ресурсы
	fmt.Println("Программа завершиена через контекст")

}

func exThree() { // с помощью waitgroup
	var wg sync.WaitGroup
	wg.Add(3) // добавляем в в счетчик 3
	go func() {
		for {
			fmt.Println("WaitGroup...")
			wg.Done() // уменьшаем счетчик
			time.Sleep(time.Second * 2)
		}

	}()

	wg.Wait() // как счетчик обнулить, программа завершиться
	fmt.Println("Программа завершиена через WG")
}

func main() {
	go exOne()
	go exTwo()
	go exThree()
	time.Sleep(5 * time.Second)
}
