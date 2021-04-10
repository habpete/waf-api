package pg_client

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type databaseConfig struct {
	DBname   string
	Port     int
	Host     string
	Username string
	Password string
}

func (dc *databaseConfig) Init(pathToConfig string) error {
	content, err := ioutil.ReadFile(pathToConfig)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(content, err)
}

type pgClient struct {
	Client *sql.database
}

func New(pathToConfig string) (IPgClient, error) {
	cfg := &databaseConfig{}
	if err := cfg.Init(pathToConfig); err != nil {
		return nil, err
	}
	return &pgClient{
		Client: nil,
	}, nil
}

func (pc *pgClient) QUERY(queryString string) (interface{}, error) {
	return nil, nil
}
