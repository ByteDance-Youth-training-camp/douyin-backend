package config

import (
	"douyin_backend/biz/mw/jwt"
	"os"
	"time"

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
	RabbitMQ struct {
		Address string `yaml:"address"`
	} `yaml:"RabbitMQ"`
	Hertz struct {
		HostPort string `yaml:"host_port"`
	} `yaml:"Hertz"`
	Jwt struct {
		Secret  string `yaml:"secret"`
		Expired int64  `yaml:"expired"`
	} `yaml:"Jwt"`
}

var Cfg Config

func InitConfig(filepath string) error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, &Cfg)
	if err != nil {
		return err
	}
	jwt.Secret = []byte(Cfg.Jwt.Secret)
	jwt.Expired = time.Minute * time.Duration(Cfg.Jwt.Expired)
	return nil
}
