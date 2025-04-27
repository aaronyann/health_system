package config

import (
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port         string        `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"read_timeout"`
		WriteTimeout time.Duration `yaml:"write_timeout"`
	} `yaml:"server"`
	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
	Blockchain struct {
		RPCURL       string `yaml:"rpc_url"`
		ContractAddr string `yaml:"contract_addr"`
		PrivateKey   string `yaml:"private_key"`
	} `yaml:"blockchain"`
	JWT struct {
		SecretKey string        `yaml:"secret_key"`
		ExpiresIn time.Duration `yaml:"expires_in"`
	} `yaml:"jwt"`
}

var AppConfig Config

func LoadConfig(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
}
