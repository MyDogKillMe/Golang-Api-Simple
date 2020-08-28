package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Welcome to the HomePage!")
}
func adminPage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Welcome to the AdminPage!")
}

func handleRequest() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/admin", adminPage)
  http.ListenAndServe(":8080", nil)
}

func main() {
  handleRequest()
}