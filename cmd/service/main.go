package main

import (
	"context"
	"log"
	"net/http"

	"homework-6/internal/message_broker"
	"homework-6/internal/service/db"
	"homework-6/internal/service/handler"
	"homework-6/internal/service/repository/studentDB"
)

const port = ":9000"

var brokers = []string{
	"127.0.0.1:9091",
	"127.0.0.1:9092",
	"127.0.0.1:9093",
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	database, err := db.NewDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB(ctx)

	studentRepo := studentDB.NewStudents(database)
	sender, err := message_broker.CreateSender(brokers, "logs")
	if err != nil {
		log.Fatal(err)
	}

	server := handler.NewServer(studentRepo, sender)

	go message_broker.Consumer(brokers)

	http.Handle("/", handler.CreateRouter(*server))
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
