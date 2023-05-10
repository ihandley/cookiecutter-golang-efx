package jobs

import (
	"midigator-portfolios/cookiecutter-golang/app/initializer"
	"midigator-portfolios/cookiecutter-golang/config"
	"midigator-portfolios/cookiecutter-golang/constants"
	"midigator-portfolios/cookiecutter-golang/instance"
	"midigator-portfolios/cookiecutter-golang/queue"
	"midigator-portfolios/cookiecutter-golang/queue/workers"
	"midigator-portfolios/cookiecutter-golang/vendors"
)

type Jobs interface {
	PingWorker() workers.PingWorker
}

type jobs struct {
	pingWorker workers.PingWorker
}

func (j *jobs) PingWorker() workers.PingWorker {
	return j.pingWorker
}

func Init(
	config config.Configuration,
	instance instance.Instance,
	services initializer.Services,
) (Jobs, error) {
	db := instance.DB()
	validation := instance.Validator()

	modelValidator := vendors.NewModelValidator(validation)

	onPingQueue, err := queue.NewQueue(
		constants.EventStream,
		constants.NatsOnPingReceived,
		config,
		instance,
	)
	if err != nil {
		return nil, err
	}
	pingWorker := workers.NewPingWorker(
		onPingQueue,
		db,
		modelValidator,
	)

	return &jobs{
		pingWorker: pingWorker,
	}, nil
}
