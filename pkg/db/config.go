package db

import "fmt"

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

func (db *Config) DNS() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User,
		db.Password, db.DBName, db.SSLMode,
	)
}
