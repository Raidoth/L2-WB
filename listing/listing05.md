Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
происходит, тоже самое что и в листинге 3
функция тест возвращает nil, но с типом customError
и при сравнивании nil без типа не равна nil с типом customError
если привести nil к customError, то будет выведено ok

```
