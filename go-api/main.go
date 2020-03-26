package api

import (
  "fmt"
  http "net/http"
  "log"

  "github.com/gorilla/mux"
)

func TodosHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("HELLO\n")
}

func RunServer() {
  r := mux.NewRouter()

  r.HandleFunc("/api/v1/todos", TodosHandler)

  srv := &http.Server{
    Handler: r,
    Addr: "127.0.0.1:8000",
  }
  log.Fatal(srv.ListenAndServe())
}
