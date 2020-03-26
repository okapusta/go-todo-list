package api

import (
  "database/sql"
  "encoding/json"
  // "fmt"
  http "net/http"
  "log"

  "github.com/gorilla/mux"
  _"github.com/go-sql-driver/mysql"
  model "github.com/okapusta/go-todo-list/go-api/model"
)

type GoApi struct {
  DB     *sql.DB
  Router *mux.Router
}

func (api *GoApi) todosHandler(w http.ResponseWriter, r *http.Request) {
  todos, err := model.GetTodos(api.DB)
  if err != nil {
    log.Fatal("Failed to get todos", err)
  }
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(todos); err != nil {
    log.Fatal("Failed to encode todos", err)
  }
}

func (api *GoApi) Initialize() {
  var err error

  api.DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/todos_development") // TODO: env var
  if err != nil {
    log.Fatal("Couldn't establish database connection")
  }
  api.Router = mux.NewRouter()
}

func (api *GoApi) RunServer() {
  api.Router.HandleFunc("/api/v1/todos", api.todosHandler).Methods("GET")

  srv := &http.Server{
    Handler: api.Router,
    Addr: "127.0.0.1:8000", // TODO: Handle timeouts and such
  }
  log.Fatal(srv.ListenAndServe())
}
