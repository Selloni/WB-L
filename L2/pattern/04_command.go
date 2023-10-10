package main

import "fmt"

/*
	КОМАНАДА Command
	Тип: Поведенческий

	Команда — это поведенческий паттерн проектирования, который превращает запросы
	в объекты, позволяя передавать их как аргументы при вызове методов, ставить запросы
	в очередь, логировать их, а также поддерживать отмену операций

	Удобно работать с сигналами передоваемымив в рограмму
==============================================================
	 Шаги реализации
	1. Создайте общий интерфейс команд и определите в нём
	метод запуска.
	2. Один за другим создайте классы конкретных команд. В каждом классе должно быть поле для хранения ссылки на один или несколько объектов-получателей, которым команда будет перенаправлять основную работу.
	Кроме этого, команда должна иметь поля для хранения параметров, которые нужны при вызове методов получателя. Значения всех этих полей команда должна получать через конструктор.
	И наконец, реализуйте основной метод команды, вызывая в нём те или иные методы получателя.
	3. Добавьте в классы отправителей поля для хранения команд. Объект-отправитель должен принимать готовый объект команды извне через конструктор, либо через сеттер команды.
	4. Измените основной код отправителей так, чтобы они делегировали выполнение действия команде.
	5. Порядок инициализации объектов должен выглядеть так:
	◦ Создаём объекты получателей.
	◦ Создаём объекты команд, связав их с получателями.
	◦ Создаём объекты отправителей, связав их с командами.
==============================================================
	Преимущества
	 Убирает прямую зависимость между объектами, вызывающими операции и объектами, которые их непосредственно выполняют.
	 Позволяет реализовать простую отмену и повтор операций.
	 Позволяет реализовать отложенный запуск команд.
	 Позволяет собирать сложные команды из простых.
	 Реализует принцип открытости/закрытости.
	Недостатки
	 Усложняет код программы за счёт дополнительных классов.
==============================================================
	Реальные приемры:
	-Графический интерфейс пользователя
	-Очереди задач
	-Транзакции базы данных
	-Управление устройствами (пульты)
*/

type CommandI interface {
	Execute() // выполнение
	Undo()    // отмена
}

// Receiver - Получатель
type Radio struct{}

func (r *Radio) On() {
	fmt.Println("О нет Джон, сегодня по сводкам обещают грозу")
}
func (r *Radio) Off() {
	fmt.Println("...psh....psh....silence")
}

// Реализуем интерфейс
type RadioOnCommand struct {
	rr Radio
}

func (r *RadioOnCommand) Execute() {
	r.rr.On()
}
func (r *RadioOnCommand) Undo() {
	r.rr.Off()
}

// Receiver - Получатель
type TV struct{}

func (r *TV) On() {
	fmt.Println("Картинка, звук, свет")
}
func (r *TV) Off() {
	fmt.Println("темный экран")
}

// Реализуем интерфейс
type TVOnCommand struct {
	rr TV
}

func (r *TVOnCommand) Execute() {
	r.rr.On() // передаем сигнал
}
func (r *TVOnCommand) Undo() {
	r.rr.Off()
}

// Invoker - инициатор
type Pult struct {
	comm CommandI
}

func (p *Pult) presButton() {
	p.comm.Execute()
}

func (p *Pult) presUndo() {
	p.comm.Undo()
}

func main() {
	tv := TVOnCommand{}
	pult := Pult{comm: &tv}
	pult.presButton()
	pult.presUndo()
	pult = Pult{comm: &RadioOnCommand{}}
	pult.presButton()
}