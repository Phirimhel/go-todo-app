package core_logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Level  string `envconfig:"LEVEL"  default:"DEBUG"`
	Folder string `envconfig:"FOLDER" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("LOGGER", &config); err != nil {
		return Config{}, fmt.Errorf("Process envconfig: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {

	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get logger config %w", err)
		fmt.Println("\033[31mPANIC: NewConfigMust config\033[0m")
		panic(err)
	}

	return config
}
