package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
	"leaderboard/models"
	"os"
)

type ServiceConfig struct {
	RedisAddr      string `env:"REDIS_ADDRESS" envDefault:"127.0.0.1"`
	Version        string `env:"VERSION" envDefault:"1.0.0"`
	HttpServerAddr string `env:"HTTP_SERVER_ADDRESS" envDefault:":8080"`
}

type Config struct {
	Symbols          map[string]models.Symbol `yaml:"symbols" json:"symbols"`
	LeaderboardDepth int                      `yaml:"leaderboardDepth" json:"leaderboardDepth"`
}

// Manager is a struct that holds the configuration for the application.
// Manager.serviceConfig holds the configuration that is loaded from the environment variables and is used to configure the service.
// Manager.config holds the configuration that is loaded from the config.yaml file and is the business configuration.
type Manager struct {
	logger        *zerolog.Logger
	serviceConfig *ServiceConfig
	config        *Config
}

// New creates a new Manager and loads the configuration from the environment variables and the config.yaml file.
func New(log *zerolog.Logger) *Manager {
	log.Info().Msg("Config Manager - Parsing Environment Variables")
	serviceCfg := ServiceConfig{}
	if err := env.Parse(&serviceCfg); err != nil {
		log.Fatal().Err(err).Stack().Msg("Failed to parse the config")
		return nil
	}
	log.Info().Msg("Config Manager - Parsing YAML Config")
	conf := &Config{}
	f, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal().Err(err).Stack().Msg("Failed to read config file")
	}
	if err := yaml.Unmarshal(f, conf); err != nil {
		log.Fatal().Err(err).Stack().Msg("Failed to parse the config file")
		return nil
	}
	log.Info().Interface("serviceConfig", &serviceCfg).Interface("config", conf).Msg("Config Manager - Config Loaded")
	return &Manager{
		logger:        log,
		serviceConfig: &serviceCfg,
		config:        conf,
	}
}

// Get returns the configuration loaded from the config.yaml file.
func (m *Manager) Get() *Config {
	return m.config
}

// GetServiceConfig returns the configuration loaded from the environment variables.
func (m *Manager) GetServiceConfig() *ServiceConfig {
	return m.serviceConfig
}
