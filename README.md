# LearnGo-Top-Ten-Words

## Task 

## Learnings
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

8. `http.Error(w, err.Error(), http.StatusInternalServerError)`
The http.Error function sends a specified HTTP response code (in this case "Internal Server Error") and error message.

## Dont understand
`t.Execute()` and `t.ExecuteTeplate()`