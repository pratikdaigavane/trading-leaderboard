package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
	"leaderboard/models"
	"os"
)

type ServiceConfig struct {
	RedisAddr string `env:"REDIS_ADDRESS" envDefault:"127.0.0.1"`
	Version   string `env:"VERSION" envDefault:"1.0.0"`
}

type Config struct {
	Symbols          map[string]models.Symbol `yaml:"symbols" json:"symbols"`
	LeaderboardDepth int                      `yaml:"leaderboardDepth" json:"leaderboardDepth"`
}

type Manager struct {
	logger        *zerolog.Logger
	serviceConfig *ServiceConfig
	config        *Config
}

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

func (m *Manager) Get() *Config {
	return m.config
}

func (m *Manager) GetServiceConfig() *ServiceConfig {
	return m.serviceConfig
}
