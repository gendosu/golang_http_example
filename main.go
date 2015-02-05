package main

import (
  "io"
  "net/http"
)

func main() {
  http.HandleFunc("/", indexHandler)
  http.ListenAndServe(":80", nil)
}

func indexHandler( w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  io.WriteString(w, "HelloWorld!\n")
}
