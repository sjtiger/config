package loader

import (
	"fmt"
	"os"

	"github.com/sjtiger/config"
	"github.com/sjtiger/config/flag"
	"github.com/sjtiger/config/log"
)

// FlagLoader loads configuration from flags.
type FlagLoader struct{}

// Load loads the command's configuration from flag arguments.
func (*FlagLoader) Load(_ []string, conf *config.Configuration) (bool, error) {
	args := os.Args[1:]
	if len(args) == 0 {
		return false, nil
	}
	if err := flag.Decode(args, conf); err != nil {
		return false, fmt.Errorf("failed to decode configuration from flags: %v", err)
	}

	log.WithoutContext().Println("Configuration loaded from flags.")

	return true, nil
}
