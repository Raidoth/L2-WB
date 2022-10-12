Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
1
2
3
4
5
6
7
8
0
0
...

функция asChan принимает значения для записи в созданный канал и возвращает его, после того как значения закончатся, канал закрывается
функция merge принимает 2 канала для записи в созданный канал и возвращает его, то есть служит для объединения 2 каналов в 1
После этого производится чтение происходит вывод из объединенного канала, но тк нет проверки на закрытие канала в функции merge, она продолжает туда писать дефолтные значения типа данных канала, в нашем случае это 0

```
