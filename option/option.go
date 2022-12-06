package option

import (
	"github.com/spf13/viper"
)

type Option func()

// WithEnvPrefix sets the environment variable prefix
func SetEnvPrefix(prefix string) Option {
	return func() {
		viper.SetEnvPrefix(prefix)
	}
}
