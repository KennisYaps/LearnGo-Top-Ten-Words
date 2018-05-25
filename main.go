package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/yappps/LearnGo-Top-Ten-Words/myLibraries"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func resultPageHandler(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("userInput")
	results := topten.GetTopTenWords(&input)

	t, err := template.ParseFiles("templates/results.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, results); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func main() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/results/", resultPageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
