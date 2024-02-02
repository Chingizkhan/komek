package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"net/url"
	"time"
)

type (
	Config struct {
		App          `yaml:"app"`
		PG           `yaml:"postgres"`
		HTTP         `yaml:"http"`
		Log          `yaml:"logger"`
		Redis        `yaml:"redis"`
		Locker       `yaml:"locker"`
		Cookie       `yaml:"cookie"`
		OauthService `yaml:"oauth_service"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	PG struct {
		User     string `env-required:"true" yaml:"user" env:"PG_USER"`
		Password string `env-required:"true" yaml:"password" env:"PG_PASSWORD"`
		Host     string `env-required:"true" yaml:"host" env:"PG_HOST"`
		Port     string `env-required:"true" yaml:"port" env:"PG_PORT"`
		Name     string `env-required:"true" yaml:"name" env:"PG_NAME"`
		PoolMax  int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		SSLMode  string `env-required:"true" yaml:"ssl_mode" env:"PG_SSL_MODE"`
	}

	HTTP struct {
		Port    string        `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		Timeout time.Duration `env-required:"true" yaml:"timeout" env:"HTTP_TIMEOUT"`
	}

	Cookie struct {
		Secret string `env-required:"true" yaml:"secret" env:"COOKIE_SECRET"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	Redis struct {
		Addr     string `env-required:"true" yaml:"addr" env:"REDIS_ADDR"`
		Password string `env-required:"true" yaml:"password" env:"REDIS_PASSWORD"`
	}

	Locker struct {
		LockTimeout time.Duration `env-required:"true" yaml:"lock_timeout" env:"LOCKER_LOCK_TIMEOUT"`
	}

	OauthService struct {
		Addr         string `env-required:"true" yaml:"addr" env:"OAUTH2_CLIENT_ADDR"`
		ClientID     string `env-required:"true" yaml:"client_id" env:"OAUTH_SERVICE_CLIENT_ID"`
		ClientSecret string `env-required:"true" yaml:"client_secret" env:"OAUTH_SERVICE_CLIENT_SECRET"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (pg *PG) DSN() string {
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(pg.User, pg.Password),
		Host:   pg.Host + ":" + pg.Port,
		Path:   pg.Name,
	}

	q := dsn.Query()

	q.Add("sslmode", pg.SSLMode)

	dsn.RawQuery = q.Encode()

	return dsn.String()
}
