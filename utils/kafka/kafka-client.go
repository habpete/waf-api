package kafka_client

type kafkaConfig struct {}

type kafkaClient struct {
	Client *kafka.Reader
}

func New () IKafkaClient {
	return &kafkaClient{}
}

func (kc *kafkaClient) Consumer (ctx context.Context) interface {
	return nil
}