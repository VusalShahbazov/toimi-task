package server

type Config struct {
	BindAddr string
	DBConfig DBConfig
}

type DBConfig struct {
	Host     string
	Username string
	Database string
	Password string
	Port     string
}

func NewConfig(bindAddr string, dbConfig DBConfig) *Config {
	return &Config{
		BindAddr: bindAddr,
		DBConfig: dbConfig,
	}
}
