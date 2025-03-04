package config

type Config struct {
	Port string
}

func New() *Config {
	return &Config{
		Port: ":8080",
	}
}
