package hw_3

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

/*Частотный анализ
Написать функцию, которая получает на вход текст и возвращает 10
самых часто встречающихся слов без учета словоформ*/

func TopStringWords(str string, topRank int) {
	pattern := regexp.MustCompile(`[^a-zA-Z ]+`)
	str = pattern.ReplaceAllString(str, "")

	words := countStringWords(
		strings.Split(
			strings.ToLower(str),
			" "))

	pairs := getPairList(words)
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})
	pairs[:10].Print()
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Print() {
	fmt.Printf("%-15v | %v\n", "Word", "Count")
	fmt.Println("=======================")
	for _, k := range p {
		fmt.Printf("%-15v | %5v\n", k.Key, k.Value)
	}
}

func countStringWords(sl []string) map[string]int {
	grouped := map[string]int{}
	for _, word := range sl {
		_, ok := grouped[word]
		if ok {
			grouped[word]++
		} else {
			grouped[word] = 1
		}
	}
	return grouped
}

func getPairList(words map[string]int) PairList {
	result := PairList{}
	for key, val := range words {
		result = append(result, Pair{key, val})
	}
	return result
}
