package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

	Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в суперклассе,
	позволяя подклассам изменять тип создаваемых объектов.

Плюсы:

	Избавляет класс от привязки к конкретным классам продуктов.
 	Выделяет код производства продуктов в одно место, упрощая поддержку кода.
 	Упрощает добавление новых продуктов в программу.
 	Реализует принцип открытости/закрытости.

Минусы:

	Может привести к созданию больших параллельных иерархий классов,
	так как для каждого класса продукта надо создать свой подкласс создателя.

*/

const (
	TServer  = "server"
	TDesktop = "desktop"
	TMobile  = "mobile"
	TLaptop  = "laptop"
)

type Computer interface {
	GetType() string
	Info()
}

func New(Type string) Computer {
	switch Type {
	case TServer:
		return NewServer()
	case TDesktop:
		return NewDesktop()
	case TMobile:
		return NewMobile()
	case TLaptop:
		return NewLaptop()
	default:
		fmt.Println("unknow type")
		return nil
	}
}

type Server struct {
	Type string
	Core int
	RAM  int
	HDD  int
}

func (s *Server) GetType() string {
	return s.Type
}

func (s *Server) Info() {
	fmt.Printf("Type: %v, Core: %d, RAM: %d Gb, HDD: %dTb\n", s.Type, s.Core, s.RAM, s.HDD)
}

func NewServer() Computer {
	return &Server{
		Type: TServer,
		Core: 20,
		RAM:  256,
		HDD:  100,
	}
}

type Desktop struct {
	Type  string
	Core  int
	RAM   int
	HDD   int
	Mouse bool
}

func (d *Desktop) GetType() string {
	return d.Type
}

func (d *Desktop) Info() {
	fmt.Printf("Type: %v, Core: %d, RAM: %d Gb, HDD: %d Tb, Mouse: %v\n", d.Type, d.Core, d.RAM, d.HDD, d.Mouse)
}

func NewDesktop() Computer {
	return &Desktop{
		Type:  TDesktop,
		Core:  12,
		RAM:   32,
		HDD:   5,
		Mouse: true,
	}
}

type Mobile struct {
	Type    string
	Core    int
	RAM     int
	HDD     int
	Buttons bool
}

func (m *Mobile) GetType() string {
	return m.Type
}

func (m *Mobile) Info() {
	fmt.Printf("Type: %v, Core: %d, RAM: %d Gb, HDD: %d Gb, Buttons: %v\n", m.Type, m.Core, m.RAM, m.HDD, m.Buttons)
}

func NewMobile() Computer {
	return &Mobile{
		Type:    TMobile,
		Core:    4,
		RAM:     8,
		HDD:     64,
		Buttons: false,
	}
}

type Laptop struct {
	Type       string
	Core       int
	RAM        int
	HDD        int
	BatteryCap int
}

func (l *Laptop) GetType() string {
	return l.Type
}

func (l *Laptop) Info() {
	fmt.Printf("Type: %v, Core: %d, RAM: %d Gb, HDD: %d Gb, Battery: %d mAh\n", l.Type, l.Core, l.RAM, l.HDD, l.BatteryCap)
}

func NewLaptop() Computer {
	return &Laptop{
		Type:       TLaptop,
		Core:       12,
		RAM:        16,
		HDD:        512,
		BatteryCap: 7000,
	}
}

var types = []string{TServer, TDesktop, TMobile, TLaptop, "notebook"}

func main() {
	var comp Computer
	for _, val := range types {
		comp = New(val)
		if comp != nil {
			comp.Info()
		}
	}

}
