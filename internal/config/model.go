package config

type Configuration struct {
	Urls             map[string]string `yaml:"urls"`
	ConnectionString string            `yaml:"connection_string"`
}
