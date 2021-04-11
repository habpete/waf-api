package kafka_client

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/segmentio/kafka-go"
	"gopkg.in/yaml.v2"
)

var kafkaErrorTemplate = "Kafka read message error: %s"

var pathToKafkaConfig string

func init() {
	pathToKafkaConfig = "../../values/kafka_values.yaml"
}

type kafkaConfig struct {
	Brokers   []string
	TopicName string
	GroupID   string
}

func (kc *kafkaConfig) Init() error {
	content, err := ioutil.ReadFile(pathToKafkaConfig)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(content, kc)
}

type kafkaClient struct {
	Client *kafka.Reader
}

func New() (IKafkaClient, error) {
	cfg := &kafkaConfig{}
	if err := cfg.Init(); err != nil {
		return nil, err
	}
	return &kafkaClient{
		Client: kafka.NewReader(kafka.ReaderConfig{
			Brokers: cfg.Brokers,
			Topic:   cfg.TopicName,
			GroupID: cfg.GroupID,
		}),
	}, nil
}

func (kc *kafkaClient) Consumer(ctx context.Context, c chan interface{}) {
	msg, err := kc.Client.ReadMessage(ctx)
	if err != nil {
		fmt.Errorf(kafkaErrorTemplate, err)
	}
	c <- msg
}
