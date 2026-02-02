package config

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	URI string
}
