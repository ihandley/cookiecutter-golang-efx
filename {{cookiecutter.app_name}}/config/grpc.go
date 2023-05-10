package config

import (
	"github.com/spf13/viper"
	"midigator-portfolios/cookiecutter-golang/logger"
)

type GrpcConfig interface {
	Port() string
}

// api holds the config for the API
type grpcConfig struct {
	env *viper.Viper
}

// Port will returns api running port
func (config *grpcConfig) Port() string {
	config.env.AutomaticEnv()
	return config.env.GetString("grpc_port")
}

func NewGrpcConfig(env *viper.Viper) GrpcConfig {
	logger.Log.Info("Reading API config...")
	return &grpcConfig{
		env: env,
	}
}
