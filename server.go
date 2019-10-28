package main

import (
  "net/http"
  "io/ioutil"
  "strings"
)

type Route struct {
  URL string
  File string
}

func m(i ...interface{}) []interface{} { return i }

func main() {
  csvBytes, _ := ioutil.ReadFile("routes.ssv")
  csvString := string(csvBytes)

  for _, line := range strings.Split(csvString, "\n") {
    parts := strings.Split(line, " ")
    if len(parts) >= 2 && parts[0][0] != '#' {
      route := &Route{ URL: parts[0], File: parts[1] }
      http.HandleFunc(route.URL, func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, route.File)
      })
    }
  }

  http.ListenAndServe(":3000", nil)
}
