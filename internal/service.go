package internal

type IServiceFX interface {}

type serviceFX struct {
	reports IReports
	kafkaClient IKafkaClient
}

func New () IServiceFX {
	return &serviceFX{}
}

