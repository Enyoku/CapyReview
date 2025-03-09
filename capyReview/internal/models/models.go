package models

type Service struct {
	URL    string  `yaml:"url"`
	Routes []Route `yaml:"routes"`
}

type Route struct {
	Path    string   `yaml:"path"`
	Target  string   `yaml:"target"`
	Methods []string `yaml:"methods"`
}
