package api

import (
  "encoding/json"
  http "net/http"
	"log"
)

func (api *GoApi) todosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := getTodos(api.DB)
	if err != nil {
		log.Fatal("Failed to get todos", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		log.Fatal("Failed to encode todos", err)
	}
}
