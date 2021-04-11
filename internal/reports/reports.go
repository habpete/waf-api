package reports

import (
	"fmt"
	"net/http"
)

type IReports interface {
	GetDataHandler(w http.ResponseWriter, req *http.Request)
}

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
	}, nil
}

func (r *reports) GetDataHandler(w http.ResponseWriter, req *http.Request) {
	result, err := r.Reader.QUERY(getDataQuery)
	if err != nil {
		fmt.Printf(err)
	}
}
