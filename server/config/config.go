package config

import "github.com/mcuadros/go-defaults"

var config *Config

type Config struct {
	Server  Server  `json:"server" yaml:"server"`
	Robust  Robust  `json:"robust" yaml:"robust"`
	Storage Storage `json:"storage" yaml:"storage"`
}

func Current(cfgs ...Cfg) *Config {
	if config == nil {
		config = New(cfgs...)
	}
	return config
}

func New(cfgs ...Cfg) *Config {
	config = new(Config)
	defaults.SetDefaults(config)

	for _, cfg := range cfgs {
		cfg(config)
	}

	return config
}

type Server struct {
	Debug       bool   `json:"debug" yaml:"debug" default:"false"`
	Prefix      string `json:"prefix" yaml:"prefix" default:"/api/v1"`
	Port        int    `json:"port" yaml:"port" default:"80"`
	GracePeriod int    `json:"gracePeriod" yaml:"gracePeriod" default:"30"`
}

type Storage struct {
	Root string `json:"root" yaml:"root" default:"/tmp/tos/data"`
}

type Robust struct {
	Retries uint `json:"retries" yaml:"retries" default:"3"`
}
