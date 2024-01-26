package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"net/url"
)

type (
	Config struct {
		App      `yaml:"app"`
		PG       `yaml:"postgres"`
		HTTP     `yaml:"http"`
		Log      `yaml:"logger"`
		KeyCloak `yaml:"keycloak"`
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
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	KeyCloak struct {
		Realm               string `env-required:"true" yaml:"realm" env:"KEYCLOAK_REALM"`
		BaseUrl             string `env-required:"true" yaml:"base_url" env:"KEYCLOAK_BASE_URL"`
		RealmRS256PublicKey string `env-required:"true" yaml:"realm_rs256_public_key" env:"KEYCLOAK_REALM_RS256_PUBLIC_KEY"`
		RestApi             struct {
			ClientId     string `env-required:"true" yaml:"client_id" env:"KEYCLOAK_REST_API_CLIENT_ID"`
			ClientSecret string `env-required:"true" yaml:"client_secret" env:"KEYCLOAK_REST_API_CLIENT_SECRET"`
		} `env-required:"true" yaml:"rest_api" env:"KEYCLOAK_REST_API"`
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
