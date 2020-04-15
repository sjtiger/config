package env

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sjtiger/config/parser"
)

// DefaultPrefix prefix of environment variable
const DefaultPrefix = ""

// Decode config
func Decode(environ []string, prefix string, element interface{}) error {
	if err := checkPrefix(prefix); err != nil {
		return err
	}

	vars := make(map[string]string)
	for _, evr := range environ {
		n := strings.SplitN(evr, "=", 2)
		if strings.HasPrefix(strings.ToUpper(n[0]), prefix) {
			key := strings.ReplaceAll(strings.ToLower(n[0]), "_", ".")
			vars[key] = n[1]
		}
	}

	rootName := strings.ToLower(prefix[:len(prefix)-1])
	return parser.Decode(vars, element, rootName)
}

func checkPrefix(prefix string) error {
	prefixPattern := `[a-zA-Z0-9]+_`
	matched, err := regexp.MatchString(prefixPattern, prefix)
	if err != nil {
		return err
	}

	if !matched {
		return fmt.Errorf("invalid prefix %q, the prefix pattern must match the following pattern: %s", prefix, prefixPattern)
	}

	return nil
}
