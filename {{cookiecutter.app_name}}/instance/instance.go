package instance

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"midigator-portfolios/cookiecutter-golang/logger"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"midigator-portfolios/cookiecutter-golang/config"
	//"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/config"
)

type Instance interface {
	Destroy() error
	DB() *pgx.Conn
	Validator() *validator.Validate
	SnsClient() *sns.Client
}

type instance struct {
	db       *pgx.Conn
	validate *validator.Validate
	sns      *sns.Client
}

// Destroy closes the connections & cleans up the instance
func (instance *instance) Destroy() error {
	return nil
}

// DB will return the instance of database
func (instance *instance) DB() *pgx.Conn {
	return instance.db
}

// Validator returns the validator
func (instance *instance) Validator() *validator.Validate {
	return instance.validate
}

// SnsClient returns the sns client
func (instance *instance) SnsClient() *sns.Client {
	return instance.sns
}

// Init initializes the instance
func Init(config config.Configuration) Instance {
	instance := &instance{}

	// Validator initialization
	instance.validate = validator.New()

	// Postgresql database configuration
	logger.Log.Info("Database connecting...")
	conn, err := pgx.Connect(context.Background(), config.PostgresConfig().ConnectionURL())
	if err != nil {
		logger.Log.Fatal(err)
	}
	defer conn.Close(context.Background())
	logger.Log.Info("Database connected successfully...")

	// SNS configuration
	logger.Log.Info("Connecting to sns...")
	cfg, err := awsconfig.LoadDefaultConfig(context.Background())
	if err != nil {
		logger.Log.Fatal(err)
	}

	cfg.EndpointResolverWithOptions = aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:       "aws",
			URL:               config.QueueConfig().SNSURL(),
			SigningRegion:     config.QueueConfig().SNSSigningRegion(),
			HostnameImmutable: true,
		}, nil
	})
	instance.sns = sns.NewFromConfig(cfg)
	logger.Log.Info("Connected to sns successfully...")

	return instance
}
