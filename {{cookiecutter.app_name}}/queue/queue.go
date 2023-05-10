package queue

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"

	//"fmt"
	"midigator-portfolios/cookiecutter-golang/config"
	"midigator-portfolios/cookiecutter-golang/instance"
	"midigator-portfolios/cookiecutter-golang/logger"
	//"strings"
)

type Message struct {
	TopicARN           string                                 `json:"topic_arn"`
	Data               []byte                                 `json:"data"`
	Attributes         map[string]types.MessageAttributeValue `json:"attributes"`
	queueMessageObject interface{}
}

//func (m *Message) Ack(ctx context.Context) error {
//	groupError := "QUEUE_MESSAGE_ACK"
//	if m.queueMessageObject == nil {
//		// do nothing
//		err := errors.New("no queue message object found to acknowledge")
//		logger.Log.WithError(err).Error(groupError)
//		return err
//	}
//
//	msgObj := m.queueMessageObject.(*)
//	err := msgObj.Ack(nats.Context(ctx))
//	if err != nil {
//		logger.Log.WithError(err).Error(groupError)
//		return err
//	}
//
//	return nil
//}

type Queue interface {
	Publish(ctx context.Context, msg *Message) error
	//Subscribe(ctx context.Context, ch chan *Message) error
	//Unsubscribe(ctx context.Context) error
	//StreamName() string
	//Subject() string
}

type queue struct {
	streamName string
	subject    string

	closeSubscription chan bool

	config   config.Configuration
	instance instance.Instance
}

func (q *queue) initialize() error {
	//groupErr := "QUEUE_INITIALIZE"
	//
	//// check if stream already exists
	//stream, err := q.jetStreamClient.StreamInfo(q.streamName)
	//if err != nil {
	//	logger.Log.WithError(err).Error(groupErr)
	//}
	//err = nil
	//if stream != nil {
	//	logger.Log.Info("Found existing Stream")
	//	q.stream = stream
	//	subjectPresent := false
	//	for _, subject := range q.stream.Config.Subjects {
	//		if subject == q.subject {
	//			subjectPresent = true
	//			logger.Log.Info("Subject found in the current stream")
	//		}
	//	}
	//	// If subject not present then add subjects to stream
	//	if !subjectPresent {
	//		logger.Log.Info("Subject not present in the stream")
	//		q.stream.Config.Subjects = append(q.stream.Config.Subjects, q.subject)
	//		q.stream, err = q.jetStreamClient.UpdateStream(&q.stream.Config)
	//		if err != nil {
	//			logger.Log.WithError(err).Error(groupErr)
	//		}
	//	}
	//	return nil
	//}

	//// If no stream present then add new stream
	//stream, err = q.jetStreamClient.AddStream(&nats.StreamConfig{
	//	Name:     q.streamName,
	//	Subjects: []string{q.subject},
	//})
	//if err != nil {
	//	logger.Log.Error("Error creating queue")
	//	logger.Log.WithError(err).Error(groupErr)
	//	return err
	//}
	//logger.Log.Infof("Stream created %v...", q.stream.Config.Name)
	//q.stream = stream
	return nil
}

func (q *queue) Publish(ctx context.Context, msg *Message) error {
	groupError := "QUEUE_PUBLISH"

	dataString := string(msg.Data)
	_, err := q.instance.SnsClient().Publish(ctx, &sns.PublishInput{
		TopicArn:          &msg.TopicARN,
		Message:           &dataString,
		MessageAttributes: msg.Attributes,
	})
	if err != nil {
		logger.Log.WithContext(ctx).WithError(err).Error(groupError)
		return err
	}

	return nil
}

//func (q *queue) Subscribe(ctx context.Context, ch chan *Message) error {
//	groupError := "QUEUE_SUBSCRIBE"
//
//	if ch == nil {
//		err := errors.New("subscription channel cannot be nil")
//		logger.Log.WithError(err).Error(groupError)
//		return err
//	}
//
//	var err error
//
//	durableName := fmt.Sprintf("%s%s", q.config.QueueConfig().NatsClientName(), strings.Replace(q.subject, ".", "_", -1))
//
//	q.subscription, err = q.jetStreamClient.Subscribe(
//		q.subject,
//		func(msg *nats.Msg) {
//			select {
//			case <-q.closeSubscription:
//				return
//			default:
//				// forward to ch
//				data := &Message{
//					Data:               msg.Data,
//					queueMessageObject: msg,
//				}
//				ch <- data
//			}
//		},
//		nats.ManualAck(),
//		nats.Durable(durableName),
//	)
//	if err != nil {
//		logger.Log.Errorf("Failed to subscribe to stream for subject %s", q.subject)
//		logger.Log.WithContext(ctx).WithError(err).Error(groupError)
//		return err
//	}
//
//	return nil
//}

//func (q *queue) Unsubscribe(ctx context.Context) error {
//	groupError := "QUEUE_UNSUBSCRIBE"
//
//	if q.subscription == nil {
//		// nothing to unsubscribe
//		return nil
//	}
//	// fire the event for closing subscription
//	q.closeSubscription <- true
//	close(q.closeSubscription)
//
//	err := q.subscription.Unsubscribe()
//	if err != nil {
//		logger.Log.WithError(err).Error(groupError)
//		return err
//	}
//
//	return nil
//}

//func (q *queue) StreamName() string {
//	return q.streamName
//}

//func (q *queue) Subject() string {
//	return q.subject
//}

// NewQueue initializes the queue client
//func NewQueue(
//	streamName string,
//	subject string,
//	config config.Configuration,
//	instance instance.Instance,
//) (Queue, error) {
//	// Create a new requst to list queues, first we will check to see if our required queue already exists
//	listQueuesRequest := sqs.ListQueuesInput{}
//
//
//	jsClient := instance.NatsJetstreamClient()
//
//	s := &queue{
//		streamName:      streamName,
//		subject:         subject,
//		jetStreamClient: jsClient,
//
//		config:   config,
//		instance: instance,
//	}
//
//	// initialize
//	err := s.initialize()
//	return s, err
//}
