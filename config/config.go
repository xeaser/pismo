package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config is the main configuration struct
type Config struct {
	Server
	AppVersion string

	// logger config
	// db config
	// common packages
}

// Server is the server configuration struct
type Server struct {
	Port string
}

// Env is the environment type
type Env string

var config *Config

const (
	local Env = "local"
	debug Env = "debug"

	configPathDefault string = "./config/config.yaml"
	configPathLocal   string = "./config/config-local.yaml"
	configPathDebug   string = "../config/config-local.yaml"
)

// Init initializes the configuration
func Init() {
	configPath := getConfigPath()
	var cfg Config
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}
	config = &cfg
}

// Get returns the current configuration
func Get() *Config {
	return config
}

// getConfigPath returns the path to the configuration file based on the environment
func getConfigPath() string {
	var configPath string
	env := Env(os.Getenv("env"))
	switch env {
	case local:
		configPath = configPathLocal
	case debug:
		configPath = configPathDebug
	default:
		configPath = configPathDefault
	}
	return configPath
}
