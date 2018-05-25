package topten

import (
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

// Category ...
type Category struct {
	Key   string
	Value int
}

var categoriesArray []Category

func convertToArray(m *map[string]int) []Category {
	for key, value := range *m {
		categoriesArray = append(categoriesArray, Category{key, value})
	}
	return categoriesArray
}
func sortBasedOnValues(arr []Category) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Value > arr[j].Value
	})
}

// GetTopTenWords ...
func GetTopTenWords(input *string) []Category {
	if len(*input) == 0 {
		return nil
	}
	text := convertToSlice(input)
	tabulate := tabulateWords(&text)
	convertToArray(&tabulate)
	sortBasedOnValues(categoriesArray)

	var TopTenWords []Category
	for index, word := range categoriesArray {
		if index < 10 {
			TopTenWords = append(TopTenWords, Category{word.Key, word.Value})
		}
	}
	return TopTenWords
}
