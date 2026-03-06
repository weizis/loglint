package analyzer

import (
	"os" 
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
    SensitiveWords []string `yaml:"sensitive_words"`
}

func LoadConfig() *Config {
    config := &Config{
        SensitiveWords: []string{
            "password",
            "token",
            "api_key",
            "secret",
            "key",
            "credential",
        },
    }

    data, err := os.ReadFile(".loglint.yaml")
    if err != nil {
        return config
    }

    err = yaml.Unmarshal(data, config)
    if err != nil {
        log.Printf("failed to parse .loglint.yaml: %v, using defaults", err)
        return config
    }

    return config
}