package core_config

import (
	"fmt"
	"os"
	"time"
)

type GlogbalConfig struct {
	TimeZone *time.Location
}

func NewGlobalConfig() (*GlogbalConfig, error) {

	tz := os.Getenv("TIME_ZONE")
	if tz == "" {
		tz = "UTC"
	}

	zone, err := time.LoadLocation(tz)
	if err != nil {
		return nil, fmt.Errorf("[global config]: load time zone,%s, %w", tz, err)
	}

	return &GlogbalConfig{
		TimeZone: zone,
	}, nil
}

func NewGlobalConfigMust() *GlogbalConfig {
	config, err := NewGlobalConfig()
	if err != nil {
		err = fmt.Errorf("[global config]: get clobal config, %w ", err)
		panic(err)
	}

	return config
}
