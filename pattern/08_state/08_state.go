package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

	Состояние — это поведенческий паттерн проектирования,
	который позволяет объектам менять поведение в зависимости от своего состояния.
	Извне создаётся впечатление, что изменился класс объекта.

Плюсы:

 	Избавляет от множества больших условных операторов машины состояний.
 	Концентрирует в одном месте код, связанный с определённым состоянием.
 	Упрощает код контекста.

Минусы:

 	Может неоправданно усложнить код, если состояний мало и они редко меняются.

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

func (h *Human) changeActivity() {
	switch h.activ.(type) {
	case *Sleep:
		h.setActivity(&Walk{})
	case *Walk:
		h.setActivity(&Work{})
	case *Work:
		h.setActivity(&Read{})
	case *Read:
		h.setActivity(&Sleep{})
	}
}

func (h *Human) justDoIt() {
	h.activ.JustDoIt()
}

func main() {
	act := Sleep{}
	hum := Human{}

	hum.setActivity(&act)

	for i := 0; i < 10; i++ {
		hum.justDoIt()
		hum.changeActivity()
	}
}
