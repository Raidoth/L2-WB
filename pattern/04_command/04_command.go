package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

	Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты,
	позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их,
	а также поддерживать отмену операций.

Плюсы:

	Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
	Позволяет реализовать простую отмену и повтор операций.
 	Позволяет реализовать отложенный запуск операций.
	Позволяет собирать сложные команды из простых.
	Реализует принцип открытости/закрытости.

Минусы:

	Усложняет код программы из-за введения множества дополнительных классов.

*/

type Command interface {
	Execute()
}

type Computer struct{}

func (c *Computer) Start() {
	fmt.Println("Computer start")
}
func (c *Computer) Restart() {
	fmt.Println("Computer restart")
}
func (c *Computer) Stop() {
	fmt.Println("Computer stop")
}

type StartCommand struct {
	Comp Computer
}

func (s *StartCommand) Execute() {
	s.Comp.Start()
}

type RestartCommand struct {
	Comp Computer
}

func (r *RestartCommand) Execute() {
	r.Comp.Restart()
}

type StopCommand struct {
	Comp Computer
}

func (s *StopCommand) Execute() {
	s.Comp.Stop()
}

type User struct {
	Start   Command
	Restart Command
	Stop    Command
}

func NewUser(start Command, restart Command, stop Command) *User {
	return &User{
		Start:   start,
		Restart: restart,
		Stop:    stop,
	}
}
func (u *User) Starts() {
	u.Start.Execute()
}

func (u *User) Stops() {
	u.Stop.Execute()
}

func (u *User) Restarts() {
	u.Restart.Execute()
}

func main() {
	computer := Computer{}
	// user := NewUser(
	// 	&StartCommand{Comp: computer},
	// 	&RestartCommand{Comp: computer},
	// 	&StopCommand{Comp: computer})
	// user.Starts()
	// user.Restarts()
	// user.Stops()

	us2 := User{Start: &StartCommand{computer},
		Restart: &RestartCommand{computer},
		Stop:    &StopCommand{}}

	us2.Starts()
	us2.Restarts()
	us2.Stops()
}
