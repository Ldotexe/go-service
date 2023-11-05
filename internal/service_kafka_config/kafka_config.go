package service_kafka_config

var brokers = []string{
	"127.0.0.1:9091",
	"127.0.0.1:9092",
	"127.0.0.1:9093",
}

var topicName = "logs"

type KafkaConfig struct {
	Brokers   []string
	TopicName string
}

func GetConfig() KafkaConfig {
	return KafkaConfig{
		Brokers:   brokers,
		TopicName: topicName,
	}
}
