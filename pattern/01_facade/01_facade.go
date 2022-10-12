package main

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern


Фасад — это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку.
Плюсы: Изолирует клиентов от компонентов сложной подсистемы.
Минусы: Фасад может привязаться ко всем классам программы


В данном случае паттерн фасад рассмотрен на примере режимов умного дома
Например: в него входят управление светом, телевизором и сигнализацией, это соответственно может быть включено, либо выключено

устанавливаем два режима для умного дома:
1) Когда пользователь уходит на работу, то выключается свет и телевизор, а сигнализация включается
2) Когда пользователь приходит домой, то выключается сигнализация и включается свет с телевизором

*/

import (
	"fmt"
	"time"
)

type TV struct {
}

func (t *TV) TvOn() {
	fmt.Println("TV is on")
}
func (t *TV) TvOff() {
	fmt.Println("TV is off")
}

type AC struct {
}

func (a *AC) AcOn() {
	fmt.Println("AC is on")
}
func (a *AC) AcOff() {
	fmt.Println("AC is off")
}

type Alarm struct {
}

func (a *Alarm) AlarmOn() {
	fmt.Println("Alarm is on")
}
func (a *Alarm) AlarmOff() {
	fmt.Println("Alarm is off")
}

type Facade struct {
	TV
	AC
	Alarm
}

func (f *Facade) GoToWork() {
	f.TV.TvOff()
	f.Alarm.AlarmOn()
	f.AC.AcOff()
}
func (f *Facade) ComeHome() {
	f.Alarm.AlarmOff()
	f.AC.AcOn()
	f.TV.TvOn()
}

func main() {
	SmartHouse := Facade{}
	SmartHouse.GoToWork()
	time.Sleep(time.Second)
	SmartHouse.ComeHome()

}
