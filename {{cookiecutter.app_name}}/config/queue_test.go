package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type QueueConfigTestSuite struct {
	suite.Suite

	mockViper *viper.Viper
	sut       QueueConfig
}

func (suite *QueueConfigTestSuite) setupConfig() {
	os.Setenv("API_BUILD_ENV", "test")
}

func (suite *QueueConfigTestSuite) SetupTest() {
	suite.setupConfig()

	suite.mockViper = viper.New()
	suite.sut = NewQueueConfig(suite.mockViper)
}

func (suite *QueueConfigTestSuite) Test_SNSSigningRegion_ShouldReturnEmpty_IfEnvNotPresent() {
	var emptySNSSigningRegion string
	os.Setenv("SNS_SIGNING_REGION", "")
	snsSigningRegion := suite.sut.SNSSigningRegion()
	suite.Equal(emptySNSSigningRegion, snsSigningRegion)
}

func (suite *QueueConfigTestSuite) Test_SNSURL_ShouldNotReturnEmpty_IfEnvPresent() {
	var envSNSURL string = "go-cookiecutter"
	os.Setenv("SNS_URL", envSNSURL)
	snsURL := suite.sut.SNSURL()
	suite.Equal(envSNSURL, snsURL)
}

func TestQueueConfig(t *testing.T) {
	suite.Run(t, &QueueConfigTestSuite{})
}
