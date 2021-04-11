package kafka_client

import "context"

type IKafkaClient interface {
	Consumer(ctx context.Context, c chan interface{})
}
