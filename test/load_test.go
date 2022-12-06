package jgconf_test

import (
	"encoding/json"
	"os"
	"testing"

	jgconf "github.com/Jinglever/go-config"
	"github.com/Jinglever/go-config/option"
	"github.com/bmizerany/assert"
)

func TestLoadYamlConfig(t *testing.T) {
	data := struct {
		Name string `mapstructure:"name"`
		Talk bool   `mapstructure:"talk"`
		Say  string `mapstructure:"say"`
	}{}

	// set env using os
	os.Setenv("JG_TALK", "false")

	err := jgconf.LoadYamlConfig("./_data/test.yml", &data, option.WithEnvPrefix("JG"))
	if err != nil {
		t.Fatal(err)
	}
	str, _ := json.Marshal(data)
	assert.Equal(t, `{"Name":"jinglever","Talk":false,"Say":"hello world"}`, string(str))
}
