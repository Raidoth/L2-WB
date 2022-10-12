package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT)
	var wg sync.WaitGroup
	buf := make([]byte, 1024)
	scanner := bufio.NewScanner(os.Stdin)
	if ln, err := net.Listen("tcp", ":16666"); err == nil {
		defer ln.Close()
		fmt.Println("Server start")
		conn, err := ln.Accept()

		if err == nil {
			go func() {
				wg.Add(1)
				defer wg.Done()
				defer fmt.Println("Write stopped")
				for {
					select {
					case <-quit:
						return
					default:
						scanner.Scan()
						input := scanner.Text()
						conn.Write([]byte(input))
					}

				}

			}()

			go func() {
				wg.Add(1)
				defer wg.Done()
				defer fmt.Println("Read stopped")
				for {
					select {
					case <-quit:
						return
					default:
						length, err := conn.Read(buf)
						if err == nil {
							fmt.Println(string(buf[0:length]))
						}
					}
				}

			}()
		}
		<-quit
		wg.Wait()
		fmt.Println("Server stopped")
		os.Exit(0)
	}

}
