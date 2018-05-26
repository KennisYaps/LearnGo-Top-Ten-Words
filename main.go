package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"

	"github.com/yappps/LearnGo-Top-Ten-Words/myLibraries"
)

// [Ignore this]: var validPath = regexp.MustCompile(`^(\/|\/results\/\?userInput=([a-zA-Z0-9+%]*))$`)

var validPath = regexp.MustCompile(`^(\/|\/results\/)$`)

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r) 
	}
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
	renderTemplate(w, "index", nil)
}
func resultPageHandler(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("userInput")
	results := topten.GetTopTenWords(&input)
	renderTemplate(w, "results", &results)
}
func main() {
	http.HandleFunc("/", makeHandler(homePageHandler))
	http.HandleFunc("/results/", makeHandler(resultPageHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
