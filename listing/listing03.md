Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false

фунция Foo возвращает интерфейс с нулевым указателем и значением типа os.PathError, при сравнении с nil, сравнение дает false,
тк у  nil значения типа нет, а у err оно равно os.PathError
если сравнивать err с nil, приведенным к os.PathError, то результат будет true
```
