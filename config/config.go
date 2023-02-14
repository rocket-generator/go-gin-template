package config

import (
	"log"
	"time"

	"github.com/jinzhu/configor"
)

// Config represents application configuration which can also be set
// via environment variable during runtime.
type Config struct {
	Environment string `default:"local" env:"ENVIRONMENT"`

	App struct {
		APP struct {
			URL       string `default:"http://localhost" env:"APP_URL"`
			Port      uint   `default:"8080" env:"APP_PORT"`
			Timezone  string `default:"UTC" env:"APP_TIMEZONE"`
			DebugMode bool   `default:"false" env:"APP_DEBUG_MODE"`
		}
		DevMode bool `default:"false" env:"DEV"`
	}

	Database struct {
		Postgres struct {
			Host     string `default:"localhost" env:"DB_HOST"`
			User     string `default:"root" env:"DB_USER"`
			Password string `default:"" env:"DB_PASSWORD"`
			Port     string `default:"5432" env:"DB_PORT"`
			Name     string `env:"DB_NAME"`
		}

		Timeout      time.Duration `default:"5" env:"DB_TIMEOUT"`
		DialTimeout  time.Duration `default:"5" env:"DB_DIAL_TIMEOUT"`
		ReadTimeout  time.Duration `default:"5" env:"DB_READ_TIMEOUT"`
		WriteTimeout time.Duration `default:"5" env:"DB_WRITE_TIMEOUT"`
	}

	Log struct {
		Level  string `env:"LOG_LEVEL" default:"info"`
		Format string `env:"LOG_FORMAT" default:"console"`
	}

	Authentication struct {
		Secret string `env:"AUTHENTICATION_SECRET" default:"default_secret"`
	}
}

// LoadConfig loads the configuration from `.env` file in the same
// directory as the application and populate the Config accordingly.
func LoadConfig() (*Config, error) {
	var config Config

	err := loadEnv()
	if err == nil {
		log.Println("Env file loaded")
	}

	err = configor.
		New(&configor.Config{AutoReload: true, AutoReloadInterval: time.Minute}).
		Load(&config)

	if err != nil {
		log.Println(err)
		log.Fatal("Error loading config")
	}

	return &config, err
}

// LoadTestConfig ...
func LoadTestConfig() (*Config, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	return config, err
}
