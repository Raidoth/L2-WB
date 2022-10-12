package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

	Цепочка Обязанностей (Chain of responsibility) - поведенческий шаблон проектирования,
	который позволяет избежать жесткой привязки отправителя запроса к получателю.
	Все возможные обработчики запроса образуют цепочку, а сам запрос перемещается по этой цепочке.
	Каждый объект в этой цепочке при получении запроса выбирает, либо закончить обработку запроса,
	либо передать запрос на обработку следующему по цепочке объекту.

Применяется:

    Когда имеется более одного объекта, который может обработать определенный запрос
    Когда надо передать запрос на выполнение одному из нескольких объект, точно не определяя, какому именно объекту
    Когда набор объектов задается динамически

Плюсы:

    Уменьшает зависимость между клиентом и обработчиками. Каждый обработчик независимо выполняет свою роль.
    Реализует принцип единственной обязанности (каждый принцип выполняет какую-то свою роль).
    Реализует принцип открытости и закрытости.

Минусы:

    Запрос может остаться необработанным.

*/

type Notifier interface {
	setNextNotifier(Notifier)
	sendMsg(string, int)
}

type SimpleNotifier struct {
	Prior int
	Next  Notifier
}

func (s *SimpleNotifier) sendMsg(mess string, prior int) {
	if prior >= s.Prior {
		fmt.Println("Simple Notifier ", mess)
	}
	if s.Next != nil {
		s.Next.sendMsg(mess, prior)
	}
}
func (s *SimpleNotifier) setNextNotifier(not Notifier) {
	s.Next = not
}

func newSimpleNotifier(prior int) *SimpleNotifier {
	return &SimpleNotifier{
		Prior: prior,
	}
}

type MiddleNotifier struct {
	Prior int
	Next  Notifier
}

func newMiddleNotifier(prior int) *MiddleNotifier {
	return &MiddleNotifier{
		Prior: prior,
	}
}

func (s *MiddleNotifier) sendMsg(mess string, prior int) {
	if prior >= s.Prior {
		fmt.Println("Middle Notifier ", mess)
	}
	if s.Next != nil {
		s.Next.sendMsg(mess, prior)
	}
}
func (s *MiddleNotifier) setNextNotifier(not Notifier) {
	s.Next = not
}

type HardNotifier struct {
	Prior int
	Next  Notifier
}

func newHardNotifier(prior int) *HardNotifier {
	return &HardNotifier{
		Prior: prior,
	}
}

func (s *HardNotifier) sendMsg(mess string, prior int) {
	if prior >= s.Prior {
		fmt.Println("Hard Notifier ", mess)
	}
	if s.Next != nil {
		s.Next.sendMsg(mess, prior)
	}
}
func (s *HardNotifier) setNextNotifier(not Notifier) {
	s.Next = not
}

func main() {

	s := newSimpleNotifier(1)
	m := newMiddleNotifier(2)
	h := newHardNotifier(3)
	s.setNextNotifier(m)
	m.setNextNotifier(h)

	s.sendMsg("Simple wrong", 1)
	fmt.Println()
	s.sendMsg("Middle wrong", 2)
	fmt.Println()
	s.sendMsg("Hard wrong", 3)
}
