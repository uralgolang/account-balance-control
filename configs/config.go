package configs

type Config struct {
	Http     string
	Postgres string
}

func NewConfig() (*Config, error) {
	//TODO read yaml and ENV
	return &Config{
		Http:     "localhost:8081",
		Postgres: "postgres://postgres:123@localhost:5432/postgres",
	}, nil
}
