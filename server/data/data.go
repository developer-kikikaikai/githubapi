package data

import (
	"github.com/kelseyhightower/envconfig"
)

type Token struct {
	Code string `json:"code"`
	Token string `json:"token"`
}

type Settings struct {
	Port string `envconfig:"server_port" default:""`
	Cert string `envconfig:"server_cert" default:""`
	Key string `envconfig:"server_key" default:""`
	ClientKey string `envconfig:"client_key" default:""`
	ClientSecret string `envconfig:"client_sercret" default:""`
}

var _settings = loadSettings()

func loadSettings() Settings {
	var settings Settings
	envconfig.Process("", &settings)
	return settings
}

func GetPort() string {
	return _settings.Port
}

func GetCert() string {
	return _settings.Cert
}

func GetKey() string {
	return _settings.Key
}

func GetClientKey() string {
	return _settings.ClientKey
}

func GetClientSecret() string {
	return _settings.ClientSecret
}
