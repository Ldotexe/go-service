package message_broker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"homework-6/internal/message_broker/kafka"
)

type Sender interface {
	SendMessage(request *http.Request, status int)
}

type operationMessage struct {
	Method  string
	Time    time.Time
	Status  int
	Success bool
}

type KafkaSender struct {
	producer *kafka.Producer
	topic    string
}

func CreateSender(brokers []string, topic string) (*KafkaSender, error) {
	producer, err := kafka.NewProducer(brokers)
	if err != nil {
		return nil, err
	}
	return NewKafkaSender(producer, topic), nil
}

func NewKafkaSender(producer *kafka.Producer, topic string) *KafkaSender {
	return &KafkaSender{
		producer: producer,
		topic:    topic,
	}
}

func (s *KafkaSender) sendAsyncMessage(message operationMessage) error {
	kafkaMsg, err := s.buildMessage(message)
	if err != nil {
		fmt.Println("Send message marshal error", err)
		return err
	}

	s.producer.SendAsyncMessage(kafkaMsg)

	fmt.Println("Send async message with key:", kafkaMsg.Key)
	return nil
}

func (s *KafkaSender) buildMessage(message operationMessage) (*sarama.ProducerMessage, error) {
	msg, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Send message marshal error", err)
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic:     s.topic,
		Value:     sarama.ByteEncoder(msg),
		Partition: -1,
		Key:       sarama.StringEncoder(message.Method),
	}, nil
}

func (s *KafkaSender) SendMessage(request *http.Request, status int) {
	err := s.sendAsyncMessage(
		operationMessage{
			Method:  request.Method,
			Time:    time.Now(),
			Status:  status,
			Success: status == http.StatusOK,
		},
	)

	if err != nil {
		fmt.Println("Send async message error: ", err)
	}
}
