package internal

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const (
	OPENAI_ENV_NAME          = "OPENAI_API_KEY"
	CONTEXT_TIMEOUT_ENV_NAME = "CONTEXT_TIMEOUT"
)

type TestConfigSuite struct {
	suite.Suite

	// The original value before running the tests
	// contained by `OPENAI_API_KEY` env
	restoreOpenaiEnvValue         string
	restoreContextTimeoutEnvValue string
}

func (s *TestConfigSuite) SetupSuite() {
	s.restoreOpenaiEnvValue = os.Getenv(OPENAI_ENV_NAME)
	s.restoreContextTimeoutEnvValue = os.Getenv(CONTEXT_TIMEOUT_ENV_NAME)
}

func (s *TestConfigSuite) TearDownSuite() {
	err := os.Setenv(OPENAI_ENV_NAME, s.restoreOpenaiEnvValue)
	if err != nil {
		s.FailNowf("can not restore %s", OPENAI_ENV_NAME)
	}
	err = os.Setenv(CONTEXT_TIMEOUT_ENV_NAME, s.restoreContextTimeoutEnvValue)
	if err != nil {
		s.FailNowf("can not restore %s", OPENAI_ENV_NAME)
	}
}

func (s *TestConfigSuite) TestGetConfig() {
	apiKey := "api-key-openai"
	contextTimeout := "3s"

	err := os.Setenv(OPENAI_ENV_NAME, apiKey)
	if err != nil {
		s.FailNow("can not set openai api key for test")
	}
	err = os.Setenv(CONTEXT_TIMEOUT_ENV_NAME, contextTimeout)
	if err != nil {
		s.FailNow("can not set context timout for test")
	}

	config, err := GetConfig()
	s.Nil(err)

	expected := &Config{
		OpenAIAPIKey:   apiKey,
		ContextTimeout: time.Duration(3) * time.Second,
	}
	fmt.Printf("expected: %v, actual : %v", expected, config)
	s.Equal(expected, config)
}

func (s *TestConfigSuite) TestValidateConfigValidConfigs() {
	apiKey := "api-key-openai"
	contextTimeout := time.Duration(3) * time.Second

	err := ValidateConfig(&Config{
		OpenAIAPIKey:   apiKey,
		ContextTimeout: contextTimeout,
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
