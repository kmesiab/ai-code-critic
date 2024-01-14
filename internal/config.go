package internal

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	goenv "github.com/Netflix/go-env"
)

const defaultContextTimeout = 5 * time.Second

// Config struct holds the configuration settings.
// Annotations for environment variables should be added here.
// Example:
//
//	type Config struct {
//	    DatabaseURL string `env:"DATABASE_URL"`
//	}
type Config struct {
	OpenAIAPIKey   string        `env:"OPENAI_API_KEY"`
	ContextTimeout time.Duration `env:"CONTEXT_TIMEOUT,default=5s"`
	// IgnoreFiles is a comma-separated list of file names to ignore.
	// We add some defaults here for now as a fallback.
	IgnoreFiles string `env:"IGNORE_FILES,default=go.mod,go.sum"`
}

// config is a singleton instance of Config
var config *Config

// once is used to ensure the config is initialized only once
var once sync.Once

// GetConfig initializes and returns the singleton instance of Config.
// It unmarshal environment variables into the Config struct and validates it.
func GetConfig() (*Config, error) {
	once.Do(func() {
		config = &Config{}

		// Unmarshal environment variables into the config struct
		_, err := goenv.UnmarshalFromEnviron(config)
		if err != nil {
			// Check for specific time duration parsing error
			if strings.Contains(err.Error(), "time: invalid duration") {
				// Set a default value for ContextTimeout
				config.ContextTimeout = defaultContextTimeout
			} else {
				Logf("Failed to unmarshal config").
					AddError(err).
					Error()

				return
			}
		}

		// Set a default value if ContextTimeout is zero (not set)
		if config.ContextTimeout == 0 {
			config.ContextTimeout = defaultContextTimeout
		}
	})

	// Validate the config after unmarshalling
	err := ValidateConfig(config)

	return config, err
}

// ValidateConfig checks if the config fields are properly set.
// It ensures that all fields are non-empty.
func ValidateConfig(config *Config) error {
	if config == nil {
		return fmt.Errorf("config is nil")
	}

	v := reflect.ValueOf(config)

	// Check if config is a pointer and dereference it
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Ensure that config is a struct
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("config is not a struct")
	}

	// Iterate over all fields of the struct
	m := v.NumField()
	for i := 0; i < m; i++ {
		// Validate that no field is empty
		if v.Field(i).String() == "" {
			return fmt.Errorf("%s must not be empty", v.Type().Field(i).Name)
		}
	}

	return nil
}
