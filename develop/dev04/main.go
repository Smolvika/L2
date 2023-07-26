package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	arr := []string{"пятак", "пятка", "пятка", "тяпка", "столик", "листок", "слиток", "моль"}
	// Приведение слова к нижнему регистру
	conversionToLower(arr)
	// Удаление дубликатов
	arr = removeDuplicates(arr)
	// Создание множеств
	res := searchAnagrams(arr)

	fmt.Println(res)

}

func searchAnagrams(s []string) map[string][]string {
	res := make(map[string][]string)

	for _, val := range s {
		word := strings.Split(val, "")
		sort.Slice(word, func(i, j int) bool {
			return word[i] < word[j]
		})

		str := strings.Join(word, "")
		res[str] = append(res[str], val)
	}

	for key, val := range res {
		delete(res, key)
		if len(val) < 2 {
			continue
		}

		res[val[0]] = val

		sort.Slice(val, func(i, j int) bool {
			return val[i] < val[j]
		})
	}

	return res
}

func conversionToLower(s []string) {
	for _, val := range s {
		val = strings.ToLower(val)
	}
}

func removeDuplicates(s []string) []string {
	check := make(map[string]struct{})
	res := make([]string, 0)

	for _, val := range s {
		if _, ok := check[val]; !ok {
			res = append(res, val)
		}
		check[val] = struct{}{}
	}

	return res
}
