package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Load() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	v := viper.New()
	v.SetConfigName("application-" + env)
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AddConfigPath("./internal/config")

	// ENV override support (important!)
	v.AutomaticEnv()
	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(
		strings.NewReplacer(".", "_"),
	)

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	for _, key := range v.AllKeys() {
		val := v.GetString(key)
		if strings.Contains(val, "${") {
			v.Set(key, os.ExpandEnv(val))
		}
	}

	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	return cfg
}
