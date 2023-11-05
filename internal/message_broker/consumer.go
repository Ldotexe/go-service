package message_broker

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/IBM/sarama"

	"homework-6/internal/message_broker/kafka"
)

type Consumer struct {
	client     sarama.ConsumerGroup
	isPaused   bool
	cancelFunc context.CancelFunc
}

func (c *Consumer) Init() *sarama.Config {
	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion

	/*
		sarama.OffsetNewest - получаем только новые сообщений, те, которые уже были игнорируются
		sarama.OffsetOldest - читаем все с самого начала
	*/
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Используется, если ваш offset "уехал" далеко и нужно пропустить невалидные сдвиги
	config.Consumer.Group.ResetInvalidOffsets = true

	// Сердцебиение консьюмера
	config.Consumer.Group.Heartbeat.Interval = 3 * time.Second

	// Таймаут сессии
	config.Consumer.Group.Session.Timeout = 60 * time.Second

	// Таймаут ребалансировки
	config.Consumer.Group.Rebalance.Timeout = 60 * time.Second

	const BalanceStrategy = "roundrobin"
	switch BalanceStrategy {
	case "sticky":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategySticky()}
	case "roundrobin":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	case "range":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	default:
		log.Printf("Unrecognized consumer group partition assignor: %s", BalanceStrategy)
	}
	return config
}

func (c *Consumer) Start(config *sarama.Config, brokers []string, topicName string) {
	log.Println("Starting a new Sarama consumer")
	/**
	 * Setup a new Sarama consumer group
	 */
	consumer := kafka.NewConsumerGroup()
	group := "service"

	ctx, cancel := context.WithCancel(context.Background())
	c.cancelFunc = cancel

	var err error
	c.client, err = sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	c.isPaused = false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := c.client.Consume(ctx, []string{topicName}, &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
		}
	}()

	<-consumer.Ready() // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	wg.Wait()

	if err = c.client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

func (c *Consumer) PauseProcessing() {
	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)
	for {
		select {
		case <-sigusr1:
			c.toggleConsumptionFlow()
		}
	}
}

func (c *Consumer) Close() {
	c.cancelFunc()
}

func (c *Consumer) toggleConsumptionFlow() {
	if c.isPaused {
		c.client.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		c.client.PauseAll()
		log.Println("Pausing consumption")
	}

	c.isPaused = !c.isPaused
}
