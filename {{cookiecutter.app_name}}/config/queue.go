package config

import "github.com/spf13/viper"

type QueueConfig interface {
	SNSPartitionID() string
	SNSURL() string
	SNSSigningRegion() string
	SNSHostnameImmutable() bool
	SNSTopicARN() string
}

// queueConfig holds the config for the queue
type queueConfig struct {
	env *viper.Viper
}

// SNSPartitionID will return SNS partition ID
func (config *queueConfig) SNSPartitionID() string {
	config.env.AutomaticEnv()
	return config.env.GetString("sns_partition_id")
}

// SNSURL will return SNS url
func (config *queueConfig) SNSURL() string {
	config.env.AutomaticEnv()
	return config.env.GetString("sns_url")
}

// SNSSigningRegion will return SNS signing region
func (config *queueConfig) SNSSigningRegion() string {
	config.env.AutomaticEnv()
	return config.env.GetString("sns_signing_region")
}

// SNSHostnameImmutable will return SNS hostname immutable
func (config *queueConfig) SNSHostnameImmutable() bool {
	config.env.AutomaticEnv()
	return config.env.GetBool("sns_hostname_immutable")
}

// SNSTopicARN will return SNS topic ARN
func (config *queueConfig) SNSTopicARN() string {
	config.env.AutomaticEnv()
	return config.env.GetString("sns_topic_arn")
}

// Address will return SNS address
func (config *queueConfig) Address() string {
	config.env.AutomaticEnv()
	return config.env.GetString("sns_address")
}

// Region will returns SNS region
func (config *queueConfig) Region() string {
	config.env.AutomaticEnv()
	config.env.SetDefault("sns_region", "us-west-2")
	return config.env.GetString("sns_region")
}

// Profile will return SNS profile
func (config *queueConfig) Profile() string {
	config.env.AutomaticEnv()
	return config.env.GetString("sns_profile")
}

// ID will return SNS ID
func (config *queueConfig) ID() string {
	config.env.AutomaticEnv()
	return config.env.GetString("sns_id")
}

// Secret will return SNS secret
func (config *queueConfig) Secret() string {
	config.env.AutomaticEnv()
	return config.env.GetString("sns_secret")
}

func NewQueueConfig(env *viper.Viper) QueueConfig {
	return &queueConfig{
		env: env,
	}
}
