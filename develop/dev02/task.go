package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)
В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func UnpackString(s string) (string, error) {
	if _, err := strconv.Atoi(s); err == nil {
		return "", errors.New("not correct string. String have only numbers")
	}

	var sb strings.Builder
	var symb rune
	const command = "\\"
	var isCommand = false
	var second = false
	for i, v := range s {
		if second {
			second = false
			continue
		}
		count, err := strconv.Atoi(string(v))
		if err != nil {
			switch {
			case string(v) == command && !isCommand:
				isCommand = true
				continue
			case string(v) == command && isCommand:
				isCommand = false
				symb = v
				sb.WriteRune(v)
				continue
			default:
				symb = v
				sb.WriteRune(symb)
			}

		} else {
			if isCommand {
				if len(s) > i+1 {
					check, err := strconv.Atoi(string(s[i+1]))
					if err == nil {
						cnt := check
						str := strconv.Itoa(count)
						for i := 0; i < cnt; i++ {
							sb.WriteString(str)
						}
						second = true
						continue
					}
				}

				sb.WriteString(fmt.Sprintf("%d", count))
				isCommand = false
				continue

			}

			for i := 0; i < count-1; i++ {
				sb.WriteRune(symb)
			}
		}
	}
	return sb.String(), nil

}
