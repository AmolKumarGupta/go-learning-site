package config

type Config struct {
	Name string
}

var App Config = Config{
	Name: "Sample site",
}
