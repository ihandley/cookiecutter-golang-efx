package runner

import (
	"context"
	"fmt"
	"midigator-portfolios/cookiecutter-golang/api/rest/router"
	"midigator-portfolios/cookiecutter-golang/app/initializer"
	"midigator-portfolios/cookiecutter-golang/config"
	"midigator-portfolios/cookiecutter-golang/instance"
	"midigator-portfolios/cookiecutter-golang/logger"
	"net/http"
	"sync"
	"time"
)

// API is the interface for rest api runner
type API interface {
	Go(ctx context.Context, wg *sync.WaitGroup)
}

type api struct {
	config   config.Configuration
	instance instance.Instance
}

func (runner *api) Go(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	apiConfig := runner.config.ApiConfig()
	logger.Log.Infof("Starting Rest API server on %v...", apiConfig.Port())
	services := initializer.Init(runner.config, runner.instance)

	routerV1 := router.Init(services)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", apiConfig.Port()),
		Handler:      routerV1,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()

}

// NewAPI returns an instance of the REST API runner
func NewAPI(config config.Configuration, instance instance.Instance) API {
	return &api{
		config:   config,
		instance: instance,
	}
}
