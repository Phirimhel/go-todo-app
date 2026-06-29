package core_auth

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Issuer    string        `envconfig:"ISSUER" required:"true"`
	JWTSecret string        `envconfig:"SECRET" required:"true"`
	TokenTTL  time.Duration `envconfig:"TOKEN_TTL" default:"30s"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("JWT", &config); err != nil {
		return Config{}, fmt.Errorf("Process JWT envconfig: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get JWT config %w", err)
		panic(err)
	}

	return config
}
