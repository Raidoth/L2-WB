package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===
Отсортировать строки (man sort)
Основное
Поддержать ключи
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
Дополнительное
Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	k := flag.Int("k", 0, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")
	b := flag.Bool("b", false, "игнорировать хвостовые пробелы")

	flag.Parse()

	filename := flag.Arg(0)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("No find file")
		os.Exit(1)
	}
	data := make([]string, 0, 3)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		data = append(data, scanner.Text())
	}

	res := Sort(data, *k, *n, *r, *u, *b)
	resFile, err := os.Create("out.txt")
	if err != nil {
		fmt.Println("Error create file")
		os.Exit(1)
	}
	for _, str := range res {
		resFile.WriteString(str + "\n")
	}
}

func Sort(text []string, k int, n, r, u, b bool) []string {
	columnSort := k
	isDefSort := true

	if b {
		tmp := make([]string, 0, len(text))
		for _, str := range text {
			tmp = append(tmp, strings.TrimSpace(str))
		}
		text = tmp

	}

	if u {
		tmp := make([]string, 0, len(text))
		isDub := false
		for i := 0; i < len(text); i++ {
			for j := i + 1; j < len(text); j++ {
				if text[i] == text[j] {
					isDub = true
					break
				}
			}
			if !isDub {
				tmp = append(tmp, text[i])
			}
			isDub = false
		}
		text = tmp
	}

	if columnSort > 0 {
		isDefSort = false
		data := make([][]string, len(text))
		for i, j := range text {
			data[i] = strings.Split(j, " ")
		}
		sort.Slice(data, func(i, j int) bool {
			var si string
			var sj string
			if len(data[i]) < columnSort {
				si = data[i][len(data[i])-1]

			} else {
				si = data[i][columnSort-1]
			}
			if len(data[j]) < columnSort {
				sj = data[j][len(data[j])-1]
			} else {
				sj = data[j][columnSort-1]
			}
			var si_lower = strings.ToLower(si)
			var sj_lower = strings.ToLower(sj)
			if si_lower == sj_lower {
				return si > sj
			}
			return si_lower < sj_lower

		})
		tmp := make([]string, 0, len(text))
		var s string
		for _, j := range data {
			s = strings.Join(j, " ")
			tmp = append(tmp, s)
		}
		text = tmp

	}

	if n {
		//сортировать по числовому значению
		isDefSort = false
		num := make([]int, 0, len(text))
		var tmpNum int
		var err error
		for _, j := range text {
			tmpNum, err = strconv.Atoi(j)
			if err != nil {
				fmt.Println("Error convert string")
				os.Exit(2)
			}
			num = append(num, tmpNum)
		}
		if r {
			sort.Sort(sort.Reverse(sort.IntSlice(num)))
		} else {
			sort.Ints(num)
		}
		tmp := make([]string, 0, len(num))
		for _, n := range num {

			tmp = append(tmp, strconv.Itoa(n))
		}
		text = tmp
	}

	if isDefSort {
		sort.Slice(text, func(i, j int) bool {
			var si string = text[i]
			var sj string = text[j]
			var si_lower = strings.ToLower(si)
			var sj_lower = strings.ToLower(sj)
			if si_lower == sj_lower {
				return si < sj
			}
			return si_lower < sj_lower
		})
	}

	if r && !n {
		sort.Sort(sort.Reverse(sort.StringSlice(text)))
	}

	return text
}
