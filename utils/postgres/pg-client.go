package pg_client

type pgClient struct {}

func New () IPgClient {
	return &pgClient{}
}

func (pc *pgClient) SELECT () (interface, error) {
	return nil, nil
}

func (pc *pgClient) INSERT () (interface, error) {
	return nil, nil
}

func (pc *pgClient) UPDATE () (interface, error) {
	return nil, nil
}

func (pc *pgClient) DELETE () (interface, error) {
	return nil, nil
}