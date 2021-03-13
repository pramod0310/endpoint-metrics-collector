package conf

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)


const (
	DefaultConfigPath = "config.yaml"
	ConfigPath        = "ENV_CONFIG_PATH"
	DefaultMetricsPath = "/metrics"
	DefaultMetricsPort = "8080"
)

type HttpEndpointConfig struct {
	TimeOut time.Duration `yaml:"TimeOut"`
	BaseURL string `yaml:"BaseURL"`
	Scheme string `yaml:"Scheme"`
	Paths []string `yaml:"Paths"`
}

type Config struct {
	MetricsPath string `yaml:"MetricsPath"`
	MetricsPort string `yaml:"MetricsPort"`
	HttpEndpointConfigs []HttpEndpointConfig `yaml:"HttpEndpointConfigs"`
}

var once sync.Once
var configInstance *Config

func NewInstance() *Config{
	once.Do(func() {
		configInstance = &Config{}
		if err := configInstance.parseAndLoadConfig(); err != nil {
			log.Fatal(err)
		}
	})

	return configInstance
}


func (c *Config) parseAndLoadConfig() error {
	configFilePath := getEnvironmentValue(DefaultConfigPath)
	configContent, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(configContent, c); err != nil {
		return err
	}

	if c.MetricsPath == "" {
		c.MetricsPath = DefaultMetricsPath
	}

	if c.MetricsPort == "" {
		c.MetricsPort = DefaultMetricsPort
	}

	fmt.Printf("%+v", c)

	return nil
}

func getEnvironmentValue(defaultValue string) string{
	if value := os.Getenv(ConfigPath); value != ""{
		return value
	}
	return defaultValue
}