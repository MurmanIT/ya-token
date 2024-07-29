package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Cloud struct {
	Token    string `yaml:"token" env-required:"true" env: "YANDEX_CLOUD_TOKEN"`
	UrlCloud string `yaml:"url" env-required:"true" enf-default:"https://iam.api.cloud.yandex.net/iam/v1/tokens"`
}

type Config struct {
	Env        string `yaml:"env" env-default:"dev" env-required:"true"`
	Cloud      `yaml:"cloud"`
	HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Port     int    `yaml:"port" env-default:"8085"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true" env: "HTTP_SERVER_PASSWORD"`
}

func LoadConfig() *Config {
	var cfg Config

	configPath := getConfig(os.Getenv("CONFIG_PATH"))

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", configPath)
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("config error: %s", err)
	}

	return &cfg
}

func getConfig(env string) string {
	if env == "" {
		log.Fatal("env is required")
	}
	return "./config/" + env + ".yaml"
}
