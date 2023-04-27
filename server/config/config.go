// Package config provides a configuration struct for the application
package config
// Config represents the application configuration
type Config struct {
    DB     *DBConfig
}
// DBConfig represents the database configuration
type DBConfig struct {
    Dialect  string // database driver
    Username string // username to connect to the database
    Password string // password to connect to the database
    Name     string // name of the database
    Charset  string // character set used by the database
}
// GetConfig returns the default configuration for the application
func GetConfig() *Config {
    return &Config{
        DB: &DBConfig{
            Dialect:  "mysql",
            Username: "root",
            Password: "password",
            Name:     "easyWay",
            Charset:  "utf8",
        },
    }
}
