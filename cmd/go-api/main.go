package main

import api "github.com/okapusta/go-todo-list/go-api"

func main()  {
  app := api.GoApi{}

  app.Initialize()
  app.RunServer()
}
