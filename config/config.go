package config

type Config struct {
	Author    string
	Publisher string
}

func LoadConfig() Config {
	cfg := Config{

		Author:    "Arvind",
		Publisher: "Your Company",
	}

	return cfg
}
