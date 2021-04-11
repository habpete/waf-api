package internal

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type serviceConfig struct {
	ROdatabaseConfig string,
	WOdatabaseConfig string,

	servicePort int
}

func (sc *serviceserviceConfig) Init () error {
	content, err := ioutil.ReadFile("")
	if err != nil {
		return err
	}
	return yaml.Unmarshal(content, sc)
}