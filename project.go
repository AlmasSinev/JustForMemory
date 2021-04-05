package main

import ("fmt"
        "net/http")

func home_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Go go go go!</h1>")
}
func help_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "It is a help page")
}

func handleReguest() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/help/", help_page)
  http.ListenAndServe(":8080", nil)
}

func main() {
  handleReguest()
}
