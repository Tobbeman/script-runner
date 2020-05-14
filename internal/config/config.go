package config

import (
	"crypto/rand"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Token string
}

const DefaultPath = "./config.yaml"

func Setup(p string) (*Config, error) {
	path := p
	if p == "" {
		path = DefaultPath
	}
	if _, err := os.Stat(path); err != nil {
		return create(path)
	} else {
		return load(path)
	}
}

func create(path string) (*Config, error) {
	c := Config{
		generateToken(10),
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
	return &c, nil
}

func generateToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}