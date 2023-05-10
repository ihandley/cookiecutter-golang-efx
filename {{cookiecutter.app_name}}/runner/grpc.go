package runner

import (
	"context"
	"fmt"
	"midigator-portfolios/cookiecutter-golang/api/grpc/router"
	"midigator-portfolios/cookiecutter-golang/app/initializer"
	"midigator-portfolios/cookiecutter-golang/config"
	"midigator-portfolios/cookiecutter-golang/instance"
	"midigator-portfolios/cookiecutter-golang/logger"
	"sync"
)

// GRPC is the interface for gRPC runner
type GRPC interface {
	Go(ctx context.Context, wg *sync.WaitGroup)
}

type grpc struct {
	config   config.Configuration
	instance instance.Instance
}

func (runner *grpc) Go(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	grpcConfig := runner.config.GrpcConfig()
	logger.Log.Infof("Starting gRPC server on %v...", grpcConfig.Port())
	services := initializer.Init(runner.config, runner.instance)

	err := router.Init(services, fmt.Sprintf(":%s", grpcConfig.Port()))
	if err != nil {
		logger.Log.WithError(err).Error("GRPC Runner")
	}
}

// NewGRPC returns an instance of the gRPC runner
func NewGRPC(config config.Configuration, instance instance.Instance) GRPC {
	return &grpc{
		config:   config,
		instance: instance,
	}
}
