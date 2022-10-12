package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

	Посетитель — это поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции,
	не изменяя классы объектов, над которыми эти операции могут выполняться.
Плюсы:

 	Упрощает добавление операций, работающих со сложными структурами объектов.
 	Объединяет родственные операции в одном классе.
 	Посетитель может накапливать состояние при обходе структуры элементов.
Минусы:

 	Паттерн не оправдан, если иерархия элементов часто меняется.
 	Может привести к нарушению инкапсуляции элементов.

*/

type Zoo struct{}
type Cinema struct{}
type Village struct{}

type IVisitor interface {
	visitZoo(zoo *Zoo)
	visitCinema(cinema *Cinema)
	visitVillage(vill *Village)
}
type Place interface {
	accept(IVisitor)
}

func (z *Zoo) accept(vis IVisitor) {
	vis.visitZoo(z)
}
func (c *Cinema) accept(vis IVisitor) {
	vis.visitCinema(c)
}
func (v *Village) accept(vis IVisitor) {
	vis.visitVillage(v)
}

type holiday struct {
	name string
}

func (h *holiday) visitZoo(z *Zoo) {
	h.name = "Zoo"
}
func (h *holiday) visitCinema(c *Cinema) {
	h.name = "Cinema"
}
func (h *holiday) visitVillage(v *Village) {
	h.name = "Village"
}

func main() {

	p := []Place{&Cinema{}, &Village{}, &Zoo{}}
	for _, place := range p {
		visitor := holiday{}
		place.accept(&visitor)
		fmt.Println(visitor.name)
	}

}
