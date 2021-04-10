package storage

type IStorage interface{}

type storage struct {
	Writer   pg_client.IPgClient
	Producer kafka_client.IKafkaClient
}

func New(pathToConfig string) (IStorage, error) {
	client, err := pg_client.New(pathToConfig)
	if err != nil {
		return nil, err
	}
	return &storage{
		Writer: client,
	}, nil
}
