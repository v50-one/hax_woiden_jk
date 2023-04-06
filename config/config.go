package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Email struct {
		Smtp string   `yaml:"smtp"`
		Port string   `yaml:"port"`
		From string   `yaml:"from"`
		Key  string   `yaml:"key"`
		To   []string `yaml:"to"`
	} `yaml:"email"`
	PushPlus struct {
		Token   string   `yaml:"token"`
		Channel []string `yaml:"channel"`
	} `yaml:"push_plus"`
}

// ReadYaml 读取配置文件
func (c *Config) ReadYaml() error {
	// 读取配置文件
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return err
	}
	return nil
}
