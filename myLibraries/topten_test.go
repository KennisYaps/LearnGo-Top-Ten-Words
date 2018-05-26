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
		"the":  2,
		"what": 1,
		"test": 1,
		"k":    1,
	}
	for key, value := range TabulateWords(&input) {
		if value != exp[key] {
			t.Errorf("Expected %d to equal %d", value, exp[key])
		}
	}
}
