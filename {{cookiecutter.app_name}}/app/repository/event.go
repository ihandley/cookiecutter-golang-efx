package repository

import (
	"context"
	"encoding/json"
	"midigator-portfolios/cookiecutter-golang/config"
	"midigator-portfolios/cookiecutter-golang/constants"
	"midigator-portfolios/cookiecutter-golang/instance"
	"midigator-portfolios/cookiecutter-golang/logger"
	"midigator-portfolios/cookiecutter-golang/queue"
)

type EventRepo interface {
	Publish(ctx context.Context, subject string, data interface{}) error
}

type eventRepo struct {
	config   config.Configuration
	instance instance.Instance
}

// Publish is used to send message to the queue
// ctx: context.Context
// subject: constants string
// data: interface{} is the data to be sent to the queue
func (repo *eventRepo) Publish(ctx context.Context, subject string, data interface{}) error {
	var (
		groupError string = "PUBLISH_EVENT"
		message    queue.Message
		err        error
		byteData   []byte
	)

	byteData, err = json.Marshal(data)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}

	message = queue.Message{
		Data: byteData,
	}

	// Prepare Queue before sending message
	queue, err := queue.NewQueue(
		constants.EventStream,
		subject,
		repo.config,
		repo.instance,
	)

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}

	// Sending message
	err = queue.Publish(ctx, &message)

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	return nil
}

func NewEventRepo(
	config config.Configuration,
	instance instance.Instance,
) EventRepo {
	return &eventRepo{
		config:   config,
		instance: instance,
	}
}
