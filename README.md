# LearnGo-Top-Ten-Words
## Task 

## Learnings
*From https://golang.org/doc/articles/wiki/#tmp_6*
1. `http.HandleFunc("/", handler)`
`http.HandleFunc` tells the http package to handle all requests to the web root ("/") with `handler`.

2. `http.ListenAndServe(":8080", nil)`
It specifying that it should listen on port 8080 on any interface (":8080"). This function will block until the program is terminated.

3. `log.Fatal(http.ListenAndServe(":8080", nil))`
`ListenAndServe` always returns an error, since it only returns when an unexpected error occurs. In order to log that error we wrap the function call with `log.Fatal `

4. `http.ResponseWriter `
An `http.ResponseWriter` value assembles the HTTP server's response; by writing to it, we send data to the HTTP client.

5. `http.Request`
`http.Request` is a data structure that represents the client HTTP request

6. `html/template` Package
Use to keep the HTML in a separate file

7. `template.ParseFiles("index.html")`
The function `template.ParseFiles` will read the contents of `index.html` and return a `*template.Template.`
It takes any number of string arguments that identify our template files, and parses those files into templates that are named after the base file name

8. `http.Error(w, err.Error(), http.StatusInternalServerError)`
The http.Error function sends a specified HTTP response code (in this case "Internal Server Error") and error message.

9. Template caching
Instead of `renderTemplate()` calling `ParseFiles` every time a page is rendered. A better approach would be to call `ParseFiles` once at program initialization, parsing all templates into a single *Template. Then we can use the `ExecuteTemplate` method to render a specific template.

10. `template.Must`
The function template.Must is a convenience wrapper that panics when passed a non-nil error value, and otherwise returns the `*Template` unaltered. A panic is appropriate here; if the templates can't be loaded the only sensible thing to do is exit the program.

11. `func (*Regexp) FindString` vs  `FindStringIndex` and `FindStringSubmatch`
If there is no match, the return value is an empty string, but it will also be empty if the regular expression successfully matches an empty string. Use `FindStringIndex` or `FindStringSubmatch` if it is necessary to distinguish these cases. 

Both `FindStringIndex` and `FindStringSubmatch` will return `nil` instead if there is no match

*From https://golang.org/pkg/regexp/#Regexp.FindString*
## Dont understand
1. `t.Execute()` and `t.ExecuteTeplate()`
2. How to get the whole url path. eg. http://localhost:8080/results/hello .`r.URL.Path` will only return /results instead of /results/hello