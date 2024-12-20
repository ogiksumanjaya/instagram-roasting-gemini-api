package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct{}

func (c *Config) Setup(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	datas := strings.Split(string(data), "\n")
	for _, env := range datas {
		e := strings.Split(env, "=")
		if len(e) >= 2 {
			os.Setenv(strings.TrimSpace(e[0]), strings.TrimSpace(strings.Join(e[1:], "=")))
		}
	}

	return nil
}

func (c *Config) GetServerAddress() string {
	restPort, err := strconv.Atoi(os.Getenv("SERVER_REST_PORT"))
	if err != nil {
		log.Println("SERVER_REST_PORT not found, using default port 8080")
	}
	return fmt.Sprintf("%s:%d", os.Getenv("SERVER_HOST"), restPort)
}

type GeminiConfig struct {
	APIKey string
}

func (c *Config) GetGeminiConfig() GeminiConfig {
	apiKey := os.Getenv("GEMINI_API_KEY")
	return GeminiConfig{
		APIKey: apiKey,
	}
}
