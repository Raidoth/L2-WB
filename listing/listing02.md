Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1
функция test выведет 2 из-за явного объявления возвращаемой переменной, что позволяет defer выполнить инкремент и вернуть измененное значение
функция anotherTest выведет 1, потому что явно не указано возвращаемое значение и инкремент произойдет уже после возвращения
```
