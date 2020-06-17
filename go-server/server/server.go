package main

import "fmt"
import "net/http"
import "log"

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<html><body>" + greeting("Code.education Rocks!") + "</body></html>")
  // fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func greeting(message string) string {
  return fmt.Sprintf("<b>%s</b>", message)
}
