package kafka_client

type IKafkaClient interface {
	Consumer() interface
}