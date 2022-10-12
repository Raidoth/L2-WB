package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*

Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем


*/

func cut(text string, _f int, _d string, _s bool) string {

	if _s {
		if !strings.Contains(text, _d) {
			return ""
		}
	}
	res := strings.Split(text, _d)

	if _f <= len(res) {
		var sb strings.Builder
		sb.WriteString(res[_f-1])
		sb.WriteString("\n")
		return sb.String()
	}

	return ""

}

func main() {
	f := flag.Int("f", 1, "fields-fields choice(columns)")
	d := flag.String("d", "\t", "delimiter - other delimiter")
	s := flag.Bool("s", false, "separated - string only delimiter")

	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	text := make([]string, 0, 3)
	isRun := true
	for isRun {
		scanner.Scan()
		line := scanner.Text()
		switch line {
		case "\\EI":
			isRun = false
		default:
			text = append(text, line)
		}
	}

	for _, v := range text {
		fmt.Print(cut(v, *f, *d, *s))
	}

}
