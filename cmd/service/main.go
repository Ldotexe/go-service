package main

import (
	"context"
	"log"
	"net/http"

	"homework-4/internal/service/db"
	"homework-4/internal/service/handler"
	"homework-4/internal/service/repository/studentDB"
)

const port = ":9000"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	database, err := db.NewDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB(ctx)

	studentRepo := studentDB.NewStudents(database)

	server := handler.Server{Repo: studentRepo}
	http.Handle("/", handler.CreateRouter(server))
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
