package helper

import (
	"log"
	"os"

	"github.com/maikwork/balanceUserAvito/internall/model"
	"gopkg.in/yaml.v3"
)

func ReadConfig(path string) model.Config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Don't find path")
	}

	cfg := model.Config{}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal("Don't unmarshall")
	}

	return cfg
}
