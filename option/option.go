package option

import (
	"strings"

	"github.com/spf13/viper"
)

type Option func()

// WithEnvPrefix sets the environment variable prefix
func WithEnvPrefix(prefix string) Option {
	return func() {
		viper.SetEnvPrefix(prefix)
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()
	}
}
