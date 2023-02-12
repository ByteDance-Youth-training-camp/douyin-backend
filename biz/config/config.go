package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MySql struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"MySql"`
	Redis struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
	} `yaml:"Redis"`
	MinIO struct {
		Endpoint        string `yaml:"endpoint"`
		AccessKeyID     string `yaml:"key"`
		SecretAccessKey string `yaml:"secret"`
	} `yaml:"MinIO"`
	Hertz struct {
		HostPort string `yaml:"host_port"`
	} `yaml:"Hertz"`
	Jwt struct {
		Secret string `yaml:"secret"`
	} `yaml:"Jwt"`
}

var Cfg Config

func InitConfig(filepath string) error {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, &Cfg)
	if err != nil {
		return err
	}
	return nil
}
