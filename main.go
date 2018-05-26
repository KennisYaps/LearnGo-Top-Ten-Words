package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/yappps/LearnGo-Top-Ten-Words/myLibraries"
)

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/results.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, c *[]topten.Category) {
	// [Qns]: why `err := templates.ExecuteTemplate(w, "templates/"+tmpl+".html", c)` is wrong
	err := templates.ExecuteTemplate(w, tmpl+".html", c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}
func resultPageHandler(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("userInput")
	results := topten.GetTopTenWords(&input)
	renderTemplate(w, "results", &results)
}
func main() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/results/", resultPageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
