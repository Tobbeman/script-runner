package config

import (
	"crypto/rand"
	"fmt"
	"gitlab.com/Tobbeman/script-runner/internal/runner"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Token            string                 `yaml:"token" validate:"required"`
	ScriptPath       string                 `yaml:"scriptPath" validate:"required"`
	Address          string                 `yaml:"address" validate:"required"`
	HrefAddress      string                 `yaml:"hrefAddress"`
	Retention        runner.RetentionConfig `ỳaml:"retention"`
	ReadTokenHeaders []string               `yaml:"readTokenHeaders" validate:"required"`
}

const DefaultConfigPath = "/etc/script-runner/config.yaml"
const DefaultScriptPath = "/etc/script-runner/scripts"
const DefaultAddress = "0.0.0.0:80"

var DefaultReadTokenHeaders = []string{
	"X-Gitlab-Token",
}

func Setup(path string) (*Config, error) {
	if _, err := os.Stat(path); err != nil {
		return create(path)
	} else {
		return load(path)
	}
}

func (c Config) HasRetention() bool {
	if c.Retention == (runner.RetentionConfig{}) {
		return true
	}
	return false
}

func create(path string) (*Config, error) {
	c := Config{
		Token:            generateToken(10),
		ScriptPath:       DefaultScriptPath,
		Address:          DefaultAddress,
		ReadTokenHeaders: DefaultReadTokenHeaders,
	}
	buf, err := yaml.Marshal(&c)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(path, buf, 0644)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func load(path string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, err
	}
	err = validator.New().Struct(&c)
	if err != nil {
		return nil, err
	}

	if c.HrefAddress == "" {
		c.HrefAddress = c.Address
	}

	return &c, nil
}

func generateToken(len int) string {
	b := make([]byte, len)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}
