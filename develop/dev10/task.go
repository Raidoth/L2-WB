package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

/*
Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Требования:
1.	Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP. После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
2.	Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
3.	При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout

*/

func main() {
	t := flag.Int("timeout", 10, "set timeout")
	flag.Parse()
	if len(os.Args) < 3 {
		os.Exit(1)
	}
	host := os.Args[len(os.Args)-2]
	port := os.Args[len(os.Args)-1]
	connect, errConn := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Duration(*t)*time.Second)
	if errConn != nil {
		fmt.Println(errConn.Error())
		os.Exit(2)
	} else {
		go func() {
			for {
				cn, errConn := net.Dial("tcp", net.JoinHostPort(host, port))
				if errConn != nil {
					fmt.Println("Connection close server down")
					connect.Close()
					os.Exit(0)
				}
				time.Sleep(1 * time.Second)
				cn.Close()
			}
		}()
		fmt.Println("Connect")
		handleConnection(connect, host, port)

	}

}
func handleConnection(conn net.Conn, host, port string) {
	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1024)
	for {
		scanner.Scan()
		input := scanner.Text()
		s, err := conn.Write([]byte(input))
		if err == nil {
			fmt.Println(s)
		}
		length, err := conn.Read(buf)
		if err == nil {
			fmt.Println(string(buf[0:length]))
		}

	}
}
