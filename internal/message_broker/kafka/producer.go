package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/pkg/errors"
)

type Producer struct {
	brokers       []string
	asyncProducer sarama.AsyncProducer
}

func newAsyncProducer(brokers []string) (sarama.AsyncProducer, error) {
	asyncProducerConfig := sarama.NewConfig()

	asyncProducerConfig.Producer.Partitioner = sarama.NewHashPartitioner
	asyncProducerConfig.Producer.RequiredAcks = sarama.WaitForAll

	asyncProducerConfig.Producer.Return.Successes = true
	asyncProducerConfig.Producer.Return.Errors = true

	asyncProducer, err := sarama.NewAsyncProducer(brokers, asyncProducerConfig)

	if err != nil {
		return nil, errors.Wrap(err, "error with async kafka-producer")
	}

	go func() {
		// Error и Retry топики можно использовать при получении ошибки
		for e := range asyncProducer.Errors() {
			fmt.Println(e.Error())
		}
	}()

	go func() {
		for m := range asyncProducer.Successes() {
			fmt.Println("Async success with key", m.Key)
		}
	}()

	return asyncProducer, nil
}

func NewProducer(brokers []string) (*Producer, error) {
	asyncProducer, err := newAsyncProducer(brokers)
	if err != nil {
		return nil, errors.Wrap(err, "error with async kafka-producer")
	}

	producer := &Producer{
		brokers:       brokers,
		asyncProducer: asyncProducer,
	}

	return producer, nil
}

func (k *Producer) SendAsyncMessage(message *sarama.ProducerMessage) {
	k.asyncProducer.Input() <- message
}

func (k *Producer) Close() error {
	// k.asyncProducer.AsyncClose() - не использовать в общем случае
	err := k.asyncProducer.Close()
	if err != nil {
		return errors.Wrap(err, "kafka.Connector.Close")
	}

	return nil
}
