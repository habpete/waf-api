package storage

import (
	"context"
	"fmt"
	kafka_client "waf/utils/kafka"
)

type IStorage interface{}

type storage struct {
	Writer     pg_client.IPgClient
	Consumer   kafka_client.IKafkaClient
	msgChannel chan interface{}
}

func New(pathToConfig string) (IStorage, error) {
	client, err := pg_client.New(pathToConfig)
	if err != nil {
		return nil, err
	}
	kafkaClient, err := kafka_client.New()
	if err != nil {
		return nil, err
	}
	return &storage{
		Writer:   client,
		Consumer: kafkaClient,
	}, nil
}

func (s *storage) StartStorage(ctx context.Context) {
	go s.Consumer.Consumer(ctx, s.msgChannel)
	go s.Write()
}

func (s *storage) Write() {
	select {
	case msg := <-s.msgChannel:
		_, err := s.Writer.QUERY(insertIntoMainQuery)
		if err != nil {
			fmt.Errorf("", msg)
			return
		}
		fmt.Printf("msg was successed inserted")
	}
}
