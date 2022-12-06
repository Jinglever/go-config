package config

import (
	"encoding/json"
	"os"
	"testing"

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

	err := LoadYamlConfig("./_test_data/test.yml", &data, option.WithEnvPrefix("JG"))
	if err != nil {
		t.Fatal(err)
	}
	str, _ := json.Marshal(data)
	t.Log(string(str))
	assert.Equal(t, `{"Name":"jinglever","Talk":false,"Say":"hello world"}`, string(str))
}
