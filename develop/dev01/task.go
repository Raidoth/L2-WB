package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===
Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	times, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		io.WriteString(os.Stderr, err.Error())
		os.Exit(1)
	} else {
		fmt.Println(time.Now())
		fmt.Println(times)
		fmt.Println(times.Date())
		fmt.Println(times.Format(time.UnixDate))
	}

}
