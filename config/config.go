package config

import (
	"encoding/json"
	"io/ioutil"
)

type ServerConfig struct {
	Type string
	Port int
}

func ParseConfig(config string) (*ServerConfig, error) {
	raw, err := ioutil.ReadFile(config)

	if err != nil {
		return nil, err
	}

	conf := ServerConfig{Port: -1}
	err = json.Unmarshal(raw, &conf)

	if err != nil {
		return nil, err
	}

	return &conf, nil
}
