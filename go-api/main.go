package api

import (
	"database/sql"
	"log"
	http "net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type GoApi struct {
	DB     *sql.DB
	Router *mux.Router
}

func (api *GoApi) Initialize() {
	var err error

	api.DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/todos_development") // TODO: env var
	if err != nil {
		log.Fatal("Couldn't establish database connection")
	}
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	api.Router = r
}

func (api *GoApi) RunServer() {
	api.Router.HandleFunc("/api/v1/todos", api.todosHandler).Methods("GET")

	srv := &http.Server{
		Handler: api.Router,
		Addr:    "127.0.0.1:8000", // TODO: Handle timeouts and such
	}
	log.Fatal(srv.ListenAndServe())
}
