package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func convertToSlice(input *string) []string {
	re := regexp.MustCompile("[^a-zA-Z0-9 ]")
	removeSpecialCharacters := re.ReplaceAllString(*input, "")
	toLowerCase := strings.ToLower(removeSpecialCharacters)
	// [Note]: .Split() return a slice
	text := strings.Split(toLowerCase, " ")
	return text
}

func tabulateWords(text *[]string) map[string]int {
	grouped := make(map[string]int)
	// [Qns]: why does `tabulate[word] = (tabulate[word] || 0) + 1` does not work`
	for _, word := range *text {
		grouped[word]++
	}
	return grouped
}

type category struct {
	Key   string
	Value int
}

var categoriesArray []category

func convertToArray(m *map[string]int) []category {
	for key, value := range *m {
		categoriesArray = append(categoriesArray, category{key, value})
	}
	return categoriesArray
}
func sortBasedOnValues(arr []category) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Value > arr[j].Value
	})
}
func getTopTenWords(input *string) []category {
	if len(*input) == 0 {
		return nil
	}
	text := convertToSlice(input)
	tabulate := tabulateWords(&text)
	convertToArray(&tabulate)
	sortBasedOnValues(categoriesArray)

	var topTenWords []category
	for index, word := range categoriesArray {
		if index < 10 {
			topTenWords = append(topTenWords, category{word.Key, word.Value})
		}
	}
	return topTenWords
}
func main() {
	input := "hi hi bye bye what a a bye f# f. es, bye! the% the fuck the erp fuck hack test lorem lorem lorem lorem lorem lorem lorem k"
	output := getTopTenWords(&input)
	fmt.Println(output, len(output))
}
