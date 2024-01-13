package internal

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const (
	OpenaiENVName         = "OPENAI_API_KEY"
	ContextTimeoutENVName = "CONTEXT_TIMEOUT"
)

type TestConfigSuite struct {
	suite.Suite

	// The original value before running the tests
	// contained by `OPENAI_API_KEY` env
	restoreOpenaiEnvValue         string
	restoreContextTimeoutEnvValue string
}

func (s *TestConfigSuite) SetupSuite() {
	s.restoreOpenaiEnvValue = os.Getenv(OpenaiENVName)
	s.restoreContextTimeoutEnvValue = os.Getenv(ContextTimeoutENVName)
}

func (s *TestConfigSuite) TearDownSuite() {
	err := os.Setenv(OpenaiENVName, s.restoreOpenaiEnvValue)
	if err != nil {
		s.FailNowf("can not restore %s", OpenaiENVName)
	}
	err = os.Setenv(ContextTimeoutENVName, s.restoreContextTimeoutEnvValue)
	if err != nil {
		s.FailNowf("can not restore %s", ContextTimeoutENVName)
	}
}

func (s *TestConfigSuite) TestGetConfig() {
	apiKey := "api-key-openai"
	contextTimeout := "5s"

	err := os.Setenv(OpenaiENVName, apiKey)
	if err != nil {
		s.FailNow("can not set openai api key for test")
	}
	err = os.Setenv(ContextTimeoutENVName, contextTimeout)
	if err != nil {
		s.FailNow("can not set context timout for test")
	}

	config, err := GetConfig()
	s.Nil(err)

	expected := &Config{
		OpenAIAPIKey:   apiKey,
		ContextTimeout: time.Duration(5) * time.Second,
		IgnoreFiles:    "go.mod",
	}
	fmt.Printf("expected: %v, actual : %v", expected, config)
	s.Equal(expected, config)
}

func (s *TestConfigSuite) TestValidateConfigValidConfigs() {
	apiKey := "api-key-openai"
	contextTimeout := time.Duration(5) * time.Second

	err := ValidateConfig(&Config{
		OpenAIAPIKey:   apiKey,
		ContextTimeout: contextTimeout,
		IgnoreFiles:    "go.mod",
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
