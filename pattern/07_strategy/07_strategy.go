package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

	Стратегия — это поведенческий паттерн проектирования,
	который определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс,
	после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

Плюсы:

 	Горячая замена алгоритмов на лету.
 	Изолирует код и данные алгоритмов от остальных классов.
 	Уход от наследования к делегированию.
 	Реализует принцип открытости/закрытости.

Минусы:

	Усложняет программу за счёт дополнительных классов.
	Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

*/

type Activity interface {
	JustDoIt()
}

type Sleep struct{}

func (s *Sleep) JustDoIt() {
	fmt.Println("Sleeping...")
}

type Walk struct{}

func (w *Walk) JustDoIt() {
	fmt.Println("Walking...")
}

type Work struct{}

func (w *Work) JustDoIt() {
	fmt.Println("Working...")
}

type Read struct{}

func (r *Read) JustDoIt() {
	fmt.Println("Reading...")
}

type Human struct {
	activ Activity
}

func (h *Human) setActivity(act Activity) {
	h.activ = act
}

func (h *Human) exeActivity() {
	h.activ.JustDoIt()
}

func main() {
	hum := Human{}

	hum.setActivity(&Sleep{})
	hum.exeActivity()

	hum.setActivity(&Walk{})
	hum.exeActivity()

	hum.setActivity(&Work{})
	hum.exeActivity()

	hum.setActivity(&Read{})
	hum.exeActivity()

	hum.setActivity(&Sleep{})
	hum.exeActivity()
}
