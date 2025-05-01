package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Services struct {
	AuthServiceAddr string `yaml:"auth_service" env:"AUTH_SERVICE" env-default:":8080"`
}

type Config struct {
	Services Services `yaml:"services"`
	// Address mean grpc server, not configuration of smtp
	Address      string `yaml:"mail_address" env:"MAIL_ADDRESS" env-default:":8080"`
	LogLevel     string `yaml:"log_level" env:"LOG_LEVEL" env-default:"DEBUG"`
	AppLevel     string `yaml:"app_level" env:"APP_LEVEL" env-default:"DEBUG"`
	TemplatePath string `yaml:"template_path" env:"TEMPLATE_PATH"`
}

func MustLoad(configPath string) Config {
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config %q: %s", configPath, err)
	}
	return cfg
}
