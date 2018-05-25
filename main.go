package main

import (
	"fmt"

	"github.com/yappps/LearnGo-Top-Ten-Words/myLibraries"
)

func main() {
	input := "hi hi bye bye what a a bye f# f. es, bye! the% the fuck the erp fuck hack test lorem lorem lorem lorem lorem lorem lorem k"
	output := topten.GetTopTenWords(&input)
	fmt.Println(output, len(output))

}
