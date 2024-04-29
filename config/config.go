package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/oauth2"
	"net/url"
	"sync"
	"time"
)

type (
	Config struct {
		App            `yaml:"app"`
		PG             `yaml:"postgres"`
		HTTP           `yaml:"http"`
		GRPC           `yaml:"grpc"`
		Log            `yaml:"logger"`
		Redis          `yaml:"redis"`
		Locker         `yaml:"locker"`
		Cookie         `yaml:"cookie"`
		Oauth2Raw      Oauth2 `yaml:"oauth2"`
		Token          `yaml:"token"`
		BankingService `yaml:"banking_service"`
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

	GRPC struct {
		Port string `yaml:"port" env-required:"true" env:"GRPC_PORT"`
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

	Oauth2 struct {
		ServiceAddr  string   `env-required:"true" yaml:"service_addr" env:"OAUTH2_SERVICE_ADDR"`
		ClientID     string   `env-required:"true" yaml:"client_id" env:"OAUTH_CLIENT_ID"`
		ClientSecret string   `env-required:"true" yaml:"client_secret" env:"OAUTH_CLIENT_SECRET"`
		AuthURL      string   `env-required:"true" yaml:"auth_url" env:"OAUTH_AUTH_URL"`
		TokenURL     string   `env-required:"true" yaml:"token_url" env:"OAUTH_TOKEN_URL"`
		RedirectURL  string   `env-required:"true" yaml:"redirect_url" env:"OAUTH_REDIRECT_URL"`
		Scopes       []string `env-required:"true" yaml:"scopes" env:"OAUTH_SCOPES"`
	}

	Token struct {
		AccessTokenLifetime  time.Duration `env-required:"true" yaml:"access_token_lifetime" env:"ACCESS_TOKEN_LIFETIME"`
		RefreshTokenLifetime time.Duration `env-required:"true" yaml:"refresh_token_lifetime" env:"REFRESH_TOKEN_LIFETIME"`
	}

	BankingService struct {
		Addr      string `env-required:"true" yaml:"addr" env:"BANKING_SERVICE_ADDR"`
		EnableTLS bool   `yaml:"enable_tls" env:"BANKING_SERVICE_ENABLE_TLS"`
	}
)

func New(url string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(url, cfg)
	if err != nil {
		return nil, fmt.Errorf("read config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("read env error: %w", err)
	}

	cfg.GetOauthConfig()

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

var (
	oauthCfg oauth2.Config
	once     sync.Once
)

func (cfg *Config) GetOauthConfig() oauth2.Config {
	once.Do(func() {
		raw := cfg.Oauth2Raw
		oauthCfg = oauth2.Config{
			ClientID:     raw.ClientID,
			ClientSecret: raw.ClientSecret,
			RedirectURL:  raw.RedirectURL,
			Scopes:       raw.Scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  raw.AuthURL,
				TokenURL: raw.TokenURL,
			},
		}
	})
	return oauthCfg
}
