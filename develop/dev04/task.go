package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

/*
Написать функцию поиска всех множеств анаграмм по словарю.

Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
1.	Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
2.	Выходные данные: ссылка на мапу множеств анаграмм
3.	Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
4.	Массив должен быть отсортирован по возрастанию.
5.	Множества из одного элемента не должны попасть в результат.
6.	Все слова должны быть приведены к нижнему регистру.
7.	В результате каждое слово должно встречаться только один раз.

*/

func Anagramm(arr []string) map[string][]string {
	cntSymb := make(map[string]map[rune]int)
	for _, j := range arr {
		j = strings.ToLower(j)
		cntSymb[j] = make(map[rune]int)
		for _, v := range j {
			cntSymb[j][v]++
		}
	}
	preResult := make(map[string][]string)
	for i, j := range cntSymb {
		for str, val := range cntSymb {
			if reflect.DeepEqual(j, val) {
				preResult[i] = append(preResult[i], str)
				delete(cntSymb, str)
			}
		}

	}
	result := make(map[string][]string)
	for _, j := range preResult {
		if len(j) > 1 {
			sort.Strings(j)
			tmp := j
			result[j[0]] = tmp
		}
	}

	return result
}

func main() {
	arr := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

	res := Anagramm(arr)
	fmt.Println(res)
}
