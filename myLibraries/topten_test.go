package topten

import (
	"testing"
)

func TestConvertToSlice(t *testing.T) {
	str := "Unit testing, (in) !Go is just"
	exp := []string{"unit", "testing", "in", "go", "is", "just"}
	for index, value := range ConvertToSlice(&str) {
		if value != exp[index] {
			t.Errorf("Expected %s to equal %s", value, exp[index])
		}
	}
}

func TestTabulateWords(t *testing.T) {
	input := []string{"hi", "hi", "bye", "bye", "what", "the", "the", "test", "k"}
	exp := map[string]int{
		"hi":   2,
		"bye":  2,
		"what": 1,
		"the":  2,
		"test": 1,
		"k":    1,
	}
	for key, value := range TabulateWords(&input) {
		if value != exp[key] {
			t.Errorf("Expected %d to equal %d", value, exp[key])
		}
	}
}

func TestAppendMapToArray(t *testing.T) {
	input := map[string]int{
		"hi":   2,
		"bye":  2,
		"what": 1,
		"the":  2,
		"test": 1,
		"k":    1,
	}
	exp := []Category{
		{"hi", 2},
		{"bye", 2},
		{"what", 1},
		{"the", 2},
		{"test", 1},
		{"k", 1},
	}
	for index, value := range AppendMapToArray(&input) {
		if value != exp[index] {
			t.Errorf("Expected %v to equal %v", value, exp[index])
		}
	}

}

func TestGetTopTenWords(t *testing.T) {
	str := "hi hi bye hi bye what bye ?what /the the hi hi hi hi hi hi test k k what o the test test test test k what o the k, jd. jd #jd test test test what k k k k i bye test i p o bye what hi"

	exp := []Category{
		{"hi", 10},
		{"test", 9},
		{"k", 8},
		{"what", 6},
		{"bye", 5},
		{"the", 4},
		{"o", 3},
		{"jd", 3},
		{"i", 2},
		{"p", 1},
	}
	for index, element := range GetTopTenWords(&str) {
		if element != exp[index] {
			t.Errorf("Expected %v to equal %v", element, exp[index])
		}
	}
}
