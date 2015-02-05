package main

import (
//  "io"
  "net/http"
  "html/template"
)

func main() {
  http.HandleFunc("/", indexHandler)
  http.ListenAndServe(":80", nil)
}

func indexHandler( w http.ResponseWriter, r *http.Request) {

  indexTemplate := template.Must(template.ParseFiles("views/foo.html"))

  indexTemplate.Execute(w, nil)
}

