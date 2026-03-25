package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Postgres PostgresConfig `mapstructure:"POSTGRES"`
	Redis    RedisConfig    `mapstructure:"REDIS"`
	Nats     NatsConfig     `mapstructure:"NATS"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"HOST"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	DB       string `mapstructure:"DB"`
	Port     int    `mapstructure:"PORT"`
	Schema   string `mapstructure:"SCHEMA"`
}

type RedisConfig struct {
	Host     string `mapstructure:"HOST"`
	Password string `mapstructure:"PASSWORD"`
	Port     int    `mapstructure:"PORT"`
}

type NatsConfig struct {
	Host        string `mapstructure:"HOST"`
	User        string `mapstructure:"USER"`
	Password    string `mapstructure:"PASSWORD"`
	Port        int    `mapstructure:"PORT"`
	MonitorPort int    `mapstructure:"MONITOR_PORT"`
}

func Load() (*Config, error) {

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// allow nested mapping like POSTGRES_HOST
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	_ = viper.ReadInConfig()

	var cfg Config

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DB,
	)
}

func (c RedisConfig) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c NatsConfig) URL() string {
	return fmt.Sprintf(
		"nats://%s:%s@%s:%d",
		c.User,
		c.Password,
		c.Host,
		c.Port,
	)
}
