package config

import "github.com/kelseyhightower/envconfig"

// Specification for basic configurations
type Specification struct {
	Port        int    `envconfig:"PORT"`
	Debug       bool   `envconfig:"DEBUG"`
	LogLevel    string `envconfig:"LOG_LEVEL" default:"info"`
	SlackToken  string `envconfig:"SLACK_TOKEN"`
	Database    Database
	Application Application
}

// Database holds the configuration for a database
type Database struct {
	DSN string `envconfig:"DATABASE_DSN"`
}

// Application represents a simple application definition
type Application struct {
	Name    string `envconfig:"APP_NAME" default:"HelloWork"`
	Version string `envconfig:"APP_VERSION" default:"1.0"`
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
