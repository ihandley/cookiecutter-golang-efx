package router

import (
	"context"
	"midigator-portfolios/cookiecutter-golang/logger"
	"midigator-portfolios/cookiecutter-golang/queue/jobs"
)

func Init(ctx context.Context, jobs jobs.Jobs, shutDownChannel chan *bool) error {
	logger.Log.Info("Initializing router")

	err := jobs.PingWorker().Invoke(ctx, shutDownChannel)
	if err != nil {
		return err
	}

	return nil
}
