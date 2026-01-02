package config

type Config struct {
	DataPath string
}

func Load() Config {
	return Config{
		DataPath: "data/tasks.json",
	}
}
