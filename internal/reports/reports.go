package reports

type IReports interface {}

type reports struct {
	Reader pg_client.IPgClient
}

func New(pathToConfig string) (IReports, error) {
	client, err := pg_client.New(pathToConfig)
	if err != nil {
		return nil, err
	}
	return &reports{
		Reader: client,
	}
}

func (r *Reports) GetDataHandler(w http.ResponseWriter, r *http.Request) {
	result, err := r.QUERY(getDataQuery)
}