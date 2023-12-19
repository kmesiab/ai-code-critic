package internal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	OPENAI_ENV_NAME = "OPENAI_API_KEY"
)

type TestConfigSuite struct {
	suite.Suite

	// The original value before running the tests
	// contained by `OPENAI_API_KEY` env
	restoreOpenaiEnvValue string
}

func (s *TestConfigSuite) SetupSuite() {
	s.restoreOpenaiEnvValue = os.Getenv(OPENAI_ENV_NAME)
}

func (s *TestConfigSuite) TearDownSuite() {
	err := os.Setenv(OPENAI_ENV_NAME, s.restoreOpenaiEnvValue)
	if err != nil {
		s.FailNowf("can not restore %s", OPENAI_ENV_NAME)
	}
}

func (s *TestConfigSuite) TestGetConfig() {
	apiKey := "api-key-openai"

	err := os.Setenv(OPENAI_ENV_NAME, apiKey)
	if err != nil {
		s.FailNow("can not set openai api key for test")
	}

	config, err := GetConfig()
	s.Nil(err)

	expected := &Config{
		OpenAIAPIKey: apiKey,
	}
	s.Equal(expected, config)
}

func (s *TestConfigSuite) TestValidateConfigValidConfigs() {
	apiKey := "api-key-openai"

	err := ValidateConfig(&Config{
		OpenAIAPIKey: apiKey,
	})

	s.Nil(err)
}

func (s *TestConfigSuite) TestValidateConfigNil() {
	err := ValidateConfig(nil)
	s.EqualError(err, "config is nil")
}

func (s *TestConfigSuite) TestValidateConfigEmptyField() {
	err := ValidateConfig(&Config{})
	s.ErrorContains(err, "must not be empty")
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(TestConfigSuite))
}
