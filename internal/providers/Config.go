package providers

import (
	"flag"
	"os"

	"github.com/kelseyhightower/envconfig"
	yaml "gopkg.in/yaml.v3"
)

var (
	configPath string
)

func loadconfig() (AppConfig, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return AppConfig{}, err
	}
	defer f.Close()
	var cfg AppConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return AppConfig{}, err
	}
	err = envconfig.Process("", &cfg)
	if err != nil {
		return AppConfig{}, err
	}
	return cfg, err

}

type AppConfig struct {
	DbConfig DbConfig `yaml:"db"`
}
type DbConfig struct {
	Host        string `yaml:"host" envconfig:"DB_HOST"`
	Port        int    `yaml:"port" envconfig:"DB_PORT"`
	User        string `yaml:"user" envconfig:"DB_USER"`
	Password    string `yaml:"password" envconfig:"DB_PASSWORD"`
	Name        string `yaml:"name" envconfig:"DB_NAME"`
	Timeout     int    `yaml:"timeout" envconfig:"DB_CONNECTION_TIMEOUT"`
	MaxOpenConn int    `yaml:"maxOpenConn" envconfig:"DB_MAX_OPEN_CONNECTIONS"`
	MaxIdleConn int    `yaml:"maxIdleConn" envconfig:"DB_MAX_OPEN_CONNECTIONS"`
	SearchPath  string `yaml:"searchPath" envconfig:"DB_SEARCH_PATH"`
}

func GetConfig(path string) (AppConfig, error) {
	flag.StringVar(&configPath, "config", path, "path to config file")
	flag.Parse()
	return loadconfig()
}
