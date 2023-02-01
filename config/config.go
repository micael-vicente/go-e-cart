package config

// Config is the struct that holds the configuration for the application
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

// DatabaseConfig is the struct that holds the configuration for the database
type DatabaseConfig struct {
	Dialect     string `yaml:"dialect"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Name        string `yaml:"name"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	AutoMigrate bool   `yaml:"autoMigrate"`
}

// ServerConfig is the struct that holds the configuration for the server
type ServerConfig struct {
	Port int `yaml:"port"`
}
