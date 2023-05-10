package initializer

import (
	"midigator-portfolios/cookiecutter-golang/app/repository"
	"midigator-portfolios/cookiecutter-golang/app/service"
	"midigator-portfolios/cookiecutter-golang/config"
	"midigator-portfolios/cookiecutter-golang/instance"
	"midigator-portfolios/cookiecutter-golang/vendors"
)

// Services is interface for all service entrypoint
type Services interface {
	TodoService() service.TodoSvc
	EventRepo() repository.EventRepo
}

type services struct {
	todoService service.TodoSvc
	eventRepo   repository.EventRepo
}

func (svc *services) TodoService() service.TodoSvc {
	return svc.todoService
}

func (svc *services) EventRepo() repository.EventRepo {
	return svc.eventRepo
}

// Init initializes services repo
func Init(config config.Configuration, instance instance.Instance) Services {
	db := instance.DB()
	validation := instance.Validator()
	modelValidator := vendors.NewModelValidator(validation)

	eventRepo := repository.NewEventRepo(config, instance)

	return &services{
		todoService: service.NewTodoSvc(
			repository.NewTodoRepo(db),
			modelValidator,
		),
		eventRepo: eventRepo,
	}
}
