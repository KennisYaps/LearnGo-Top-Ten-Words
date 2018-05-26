package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"regexp"

	"github.com/yappps/LearnGo-Top-Ten-Words/myLibraries"
)

// var validPath = regexp.MustCompile(`^(\/|\/results\/\?userInput=([a-zA-Z0-9+%]*))$`)
var validPath = regexp.MustCompile(`^(\/|\/results\/)$`)

func validatePath(w http.ResponseWriter, r *http.Request) error {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return errors.New("Invalid Page")
	}
	return nil 
}

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/results.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, c *[]topten.Category) {
	// [Qns]: why `err := templates.ExecuteTemplate(w, "templates/"+tmpl+".html", c)` is wrong
	err := templates.ExecuteTemplate(w, tmpl+".html", c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	err := validatePath(w, r)
	if err != nil {
		return
	}
	renderTemplate(w, "index", nil)
}
func resultPageHandler(w http.ResponseWriter, r *http.Request) {
	err := validatePath(w, r)
	if err != nil {
		return
	}
	input := r.FormValue("userInput")
	results := topten.GetTopTenWords(&input)
	renderTemplate(w, "results", &results)
}
func main() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/results/", resultPageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
