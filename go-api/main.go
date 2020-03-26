package api

import (
  "fmt"
  http "net/http"
  "log"

  "github.com/gorilla/mux"
)

type GoApi struct {
  Router *mux.Router
}

func (api *GoApi) todosHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("HELLO\n")
}

func (api *GoApi) Initialize() {
  api.Router = mux.NewRouter() // TODO: Database connection
}

func (api *GoApi) RunServer() {
  api.Router.HandleFunc("/api/v1/todos", api.todosHandler).Methods("GET")

  srv := &http.Server{
    Handler: api.Router,
    Addr: "127.0.0.1:8000", // TODO: Handle timeouts and such
  }
  log.Fatal(srv.ListenAndServe())
}
