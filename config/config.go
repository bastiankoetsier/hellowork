package config

import "github.com/kelseyhightower/envconfig"

// Specification for basic configurations
type Specification struct {
	LogLevel   string `envconfig:"LOG_LEVEL" default:"info"`
	SlackToken string `envconfig:"SLACK_TOKEN" required:"true"`
}

//LoadEnv loads environment variables
func LoadEnv() (*Specification, error) {
	var config Specification
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
