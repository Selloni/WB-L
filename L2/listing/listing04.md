Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
значения 0-9
а после deadlock - так как канал ждет записи в себя, которой не случится
```
что бы избежать ошибки нужно выполнить close(ch)