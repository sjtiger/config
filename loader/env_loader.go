package loader

import (
	"fmt"
	"os"
	"strings"

	"github.com/sjtiger/config"
	"github.com/sjtiger/config/env"
	"github.com/sjtiger/config/log"
)

// EnvLoader loads configurations from environment variables
type EnvLoader struct{}

// Load config
func (e *EnvLoader) Load(_ []string, conf *config.Configuration) (bool, error) {
	vars := env.FindPrefixedEnvVars(os.Environ(), env.DefaultPrefix, conf)
	if len(vars) == 0 {
		return false, nil
	}

	if err := env.Decode(vars, env.DefaultPrefix, conf); err != nil {
		log.WithoutContext().Debug("environment variables", strings.Join(vars, ", "))
		return false, fmt.Errorf("failed to decode configuration from environment variables: %v ", err)
	}

	log.WithoutContext().Println("Configuration loaded from environment variables.")

	return true, nil
}
