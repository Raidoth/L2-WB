package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

Плюсы:

 	Позволяет создавать продукты пошагово.
 	Позволяет использовать один и тот же код для создания различных продуктов.
 	Изолирует сложный код сборки продукта от его основной бизнес-логики.

Минусы:

 	Усложняет код программы из-за введения дополнительных классов.
 	Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.

*/

type Computer struct {
	Core int
	RAM  int
	HDD  int
}

func (c *Computer) Info() {
	fmt.Printf("Core: %d, RAM: %d Gb, HDD: %d Gb\n", c.Core, c.RAM, c.HDD)
}

type IComputerBuilder interface {
	Core(int) IComputerBuilder
	RAM(int) IComputerBuilder
	HDD(int) IComputerBuilder
	Build() Computer
}
type ComputerBuilder struct {
	core int
	ram  int
	hdd  int
}

func NewComputerBuilder() IComputerBuilder {
	return &ComputerBuilder{}
}

func (c *ComputerBuilder) Core(val int) IComputerBuilder {
	c.core = val
	return c
}

func (c *ComputerBuilder) RAM(val int) IComputerBuilder {
	c.ram = val
	return c
}
func (c *ComputerBuilder) HDD(val int) IComputerBuilder {
	c.hdd = val
	return c
}

func (c *ComputerBuilder) Build() Computer {
	return Computer{
		Core: c.core,
		RAM:  c.ram,
		HDD:  c.hdd,
	}
}

type Desktop struct {
	ComputerBuilder
}

func (d *Desktop) Build() Computer {
	return Computer{
		Core: 12,
		RAM:  16,
		HDD:  5,
	}
}

func NewDesktop() IComputerBuilder {
	return &Desktop{}
}

type NoteBook struct {
	ComputerBuilder
}

func (d *NoteBook) Build() Computer {
	return Computer{
		Core: 12,
		RAM:  16,
		HDD:  1,
	}
}

func NewNotebook() IComputerBuilder {
	return &NoteBook{}
}

func main() {

	desktopBuilder := NewDesktop()
	desktop := desktopBuilder.Build()
	desktop.Info()

	notebookBuilder := NewNotebook()
	notebook := notebookBuilder.Build()
	notebook.Info()

	compBuilder := NewComputerBuilder()
	comp := compBuilder.Core(5).RAM(2).HDD(3).Build()
	comp.Info()

}
