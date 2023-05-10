package config

import (
	"fmt"
	"midigator-portfolios/cookiecutter-golang/logger"

	"github.com/spf13/viper"
)

type PostgresConfig interface {
	Host() string
	Port() string
	Database() string
	User() string
	Password() string
	ConnectionURL() string
	SSLMode() string
}

// postgres is the config object for postgres database
type postgresConfig struct {
	env *viper.Viper
}

// Host returns database hostname
func (config *postgresConfig) Host() string {
	config.env.AutomaticEnv()
	return config.env.GetString("postgres_host")
}

// Port returns database port
func (config *postgresConfig) Port() string {
	config.env.AutomaticEnv()
	return config.env.GetString("postgres_port")
}

// Database returns database name
func (config *postgresConfig) Database() string {
	config.env.AutomaticEnv()
	return config.env.GetString("postgres_db")
}

// User returns database username
func (config *postgresConfig) User() string {
	config.env.AutomaticEnv()
	return config.env.GetString("postgres_user")
}

// Password returns database password
func (config *postgresConfig) Password() string {
	config.env.AutomaticEnv()
	return config.env.GetString("postgres_password")
}

// SSLMode returns database SSL Mode
func (config *postgresConfig) SSLMode() string {
	config.env.AutomaticEnv()
	return config.env.GetString("postgres_ssl_mode")
}

// ConnectionURL returns connection url for postgresql database
func (config *postgresConfig) ConnectionURL() string {
	url := config.env.GetString("postgres_url")
	if len(url) > 0 {
		return url
	}
	return fmt.Sprintf(`postgres://%v:%v@%v:%v/%v?sslmode=%v`, config.User(), config.Password(), config.Host(), config.Port(), config.Database(), config.SSLMode())
}

func NewPostgresConfig(env *viper.Viper) PostgresConfig {
	logger.Log.Info("Reading postgresql database configuration...")
	return &postgresConfig{
		env: env,
	}
}
