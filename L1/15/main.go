//К каким негативным последствиям может привести данный
//фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.
//
// не эфективное использование рурсов, при условии что нам нужно определнное количесво
// байт, не следует выделать слишком большой размер, так как они буду висеть при всей работе программы
//var justString string
//func someFunc() {
//создаеться большое представление
//	v := createHugeString(1 << 10)
// происходить копирование первых 100 символов
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

package main

func createHugeString(len int) string {
	var str string
	for i := 0; i < len; i++ {
		// заполнение строки
	}
	return str
}

func main() {
	createHugeString(100)
}
