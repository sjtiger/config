package loader

import (
	"github.com/sjtiger/config"
	"github.com/sjtiger/config/log"
)

// ResourceLoader is a configuration resource loader.
type ResourceLoader interface {
	// Load populates cmd.Configuration, optionally using args to do so.
	Load(args []string, conf *config.Configuration) (bool, error)
}

// LoadConfig set all configuration
func LoadConfig(loaders []ResourceLoader, conf *config.Configuration) {
	for _, loader := range loaders {
		_, err := loader.Load([]string{}, conf)
		if err != nil {
			log.WithoutContext().Errorf("error load config: %v", err)
		}
	}
}
