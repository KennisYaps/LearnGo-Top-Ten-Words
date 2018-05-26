package topten

import (
	"testing"
)

func TestConvertToSlice(t *testing.T) {
	input := "Unit testing, (in) !Go is just"
	exp := []string{"unit", "testing", "in", "go", "is", "just"}
	for index, value := range ConvertToSlice(&input) {
		if value != exp[index] {
			t.Errorf("Expected %s to equal %s", value, exp[index])
		}
	}
}

