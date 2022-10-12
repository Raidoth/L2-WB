package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).

Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки

*/

func main() {
	A := flag.Int("A", 0, "\"after\" печатать +N строк после совпадения")
	B := flag.Int("B", 0, "\"before\" печатать +N строк до совпадения")
	C := flag.Int("C", 0, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	c := flag.Bool("c", false, "\"count\" (количество строк)")
	i := flag.Bool("i", false, "\"ignore-case\" (игнорировать регистр)")
	v := flag.Bool("v", false, "\"invert\" (вместо совпадения, исключать)")
	F := flag.Bool("F", false, "\"fixed\", точное совпадение со строкой, не паттерн")
	n := flag.Bool("n", false, "\"line num\", напечатать номер строки")
	flag.Parse()
	filename := flag.Arg(1)
	find := flag.Arg(0)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("No find file")
		os.Exit(1)
	}
	data := make([]string, 0, 5)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		data = append(data, scanner.Text())
	}

	grep(data, find, *A, *B, *C, *c, *i, *v, *F, *n)
}

func grep(data []string, find string, A, B, C int, c, i, v, F, n bool) {

	if c {
		count := 0
		if i {
			for _, val := range data {
				if strings.Contains(strings.ToLower(val), strings.ToLower(find)) {
					count++
				}
			}
		} else {

			for _, val := range data {
				if strings.Contains(val, find) {
					count++
				}
			}
		}
		fmt.Println(count)
		return
	}

	result := make([]string, 0, 5)
	if F {
		if n {
			for i, val := range data {
				ss := strconv.Itoa(i)
				if val == find {
					result = append(result, ss+":"+val)
				}
			}
		} else {
			for _, val := range data {
				if val == find {
					result = append(result, val)
				}
			}
		}

		for _, v := range result {
			fmt.Println(v)
		}
		return
	}
	if v {
		if i {
			for _, val := range data {
				if !strings.Contains(strings.ToLower(val), strings.ToLower(find)) {
					result = append(result, val)
				}
			}
		} else {
			for _, val := range data {
				if !strings.Contains(val, find) {
					result = append(result, val)
				}
			}
		}

		for _, v := range result {
			fmt.Println(v)
		}
		return
	}

	if C > 0 {
		if i {

			for i, val := range data {
				last := i + C + 1
				start := i - C
				if strings.Contains(strings.ToLower(val), strings.ToLower(find)) {
					if i != 0 {
						if !strings.Contains(strings.ToLower(data[i-1]), strings.ToLower(find)) {

							if start < 0 {
								for s := 0; s != i; s++ {
									result = append(result, data[s])
								}
							} else {
								for s := start; s != i; s++ {
									result = append(result, data[s])
								}
							}
						}
					}

					if i < len(data)-1 {
						if !strings.Contains(strings.ToLower(data[i+1]), strings.ToLower(find)) {

							if last > len(data) {
								for s := i; s < len(data); s++ {
									result = append(result, data[s])
								}
							} else {
								for s := i; s < last; s++ {

									result = append(result, data[s])
								}
							}
							continue
						}
					}

					result = append(result, val)

				}
			}

		} else {
			addsArray := make([]int, 0, 5)
			for i, val := range data {
				last := i + C + 1
				start := i - C
				if strings.Contains(val, find) {
					if i != 0 {
						if !strings.Contains(data[i-1], find) {

							if start < 0 {
								for s := 0; s != i; s++ {
									if data[s] == find {
										break
									}
									addsArray = append(addsArray, s)
									result = append(result, data[s])
								}
							} else {
								for s := start; s != i; s++ {
									if data[s] == find {
										break
									}
									addsArray = append(addsArray, s)
									result = append(result, data[s])
								}
							}
						}
					}

					if last > start {
						if i < len(data)-1 {
							if !strings.Contains(data[i+1], find) {

								if last > len(data) {
									for s := i; s < len(data); s++ {
										addsArray = append(addsArray, s)
										result = append(result, data[s])
									}
								} else {
									for s := i; s < last; s++ {
										addsArray = append(addsArray, s)
										result = append(result, data[s])
									}
								}
								continue
							}
						}
					}
					addsArray = append(addsArray, i)
					result = append(result, val)

				}
			}

			for _, val := range addsArray {
				fmt.Println(val)
			}

		}
	} else if A > 0 {
		if i {
			for i, val := range data {
				if strings.Contains(strings.ToLower(val), strings.ToLower(find)) {
					if i != 0 {
						if !strings.Contains(strings.ToLower(data[i-1]), strings.ToLower(find)) {
							start := i - A
							if start < 0 {
								for s := 0; s != i; s++ {
									result = append(result, data[s])
								}
							} else {
								for s := start; s != i; s++ {
									result = append(result, data[s])
								}
							}
						}
					}

					result = append(result, val)

				}
			}
		} else {

			for i, val := range data {
				if strings.Contains(val, find) {
					if i != 0 {
						if !strings.Contains(data[i-1], find) {
							start := i - A
							if start < 0 {
								for s := 0; s != i; s++ {
									result = append(result, data[s])
								}
							} else {
								for s := start; s != i; s++ {
									result = append(result, data[s])
								}
							}
						}
					}

					result = append(result, val)

				}
			}
		}

	} else if B > 0 {
		if i {
			for i, val := range data {
				if strings.Contains(strings.ToLower(val), strings.ToLower(find)) {
					if i < len(data)-1 {
						if !strings.Contains(strings.ToLower(data[i+1]), strings.ToLower(find)) {
							last := i + B + 1
							if last > len(data) {
								for s := i; s < len(data); s++ {
									result = append(result, data[s])
								}
							} else {
								for s := i; s < last; s++ {
									result = append(result, data[s])
								}
							}
							continue
						}
					}

					result = append(result, val)

				}
			}
		} else {
			for i, val := range data {
				if strings.Contains(val, find) {
					if i < len(data)-1 {
						if !strings.Contains(data[i+1], find) {
							last := i + B + 1
							if last > len(data) {
								for s := i; s < len(data); s++ {
									result = append(result, data[s])
								}
							} else {
								for s := i; s < last; s++ {
									result = append(result, data[s])
								}
							}
							continue
						}
					}

					result = append(result, val)

				}
			}
		}
	} else {
		if i {
			if n {
				for i, val := range data {
					if strings.Contains(strings.ToLower(val), find) {
						ss := strconv.Itoa(i)
						result = append(result, ss+":"+val)
					}
				}
			} else {
				for _, val := range data {
					if strings.Contains(strings.ToLower(val), find) {
						result = append(result, val)
					}
				}
			}
		} else {
			if n {
				for i, val := range data {
					ss := strconv.Itoa(i)
					if strings.Contains(val, find) {
						result = append(result, ss+":"+val)
					}
				}
			} else {
				for _, val := range data {

					if strings.Contains(val, find) {
						result = append(result, val)
					}
				}
			}
		}
	}

	for _, v := range result {
		fmt.Println(v)
	}

}
