package main

import (
	"context"
	"log"
	"net/http"

	"homework-6/internal/message_broker"
	"homework-6/internal/service/db"
	"homework-6/internal/service/handler"
	"homework-6/internal/service/repository/studentDB"
	"homework-6/internal/service_kafka_config"
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

	kafkaConfig := service_kafka_config.GetConfig()

	studentRepo := studentDB.NewStudents(database)
	sender, err := message_broker.CreateSender(kafkaConfig.Brokers, kafkaConfig.TopicName)
	if err != nil {
		log.Fatal(err)
	}

	server := handler.NewServer(studentRepo, sender)

	var consumer message_broker.Consumer
	config := consumer.Init()
	go consumer.Start(config, kafkaConfig.Brokers, kafkaConfig.TopicName)
	go consumer.PauseProcessing()
	defer consumer.Close()

	http.Handle("/", handler.CreateRouter(*server))
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
