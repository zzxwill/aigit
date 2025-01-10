package llm

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Provider struct {
	APIKey   string `json:"api_key"`
	Endpoint string `json:"endpoint"`
}

type Config struct {
	CurrentProvider string              `json:"current_provider"`
	Providers       map[string]Provider `json:"providers"`
}

func NewConfig() *Config {
	return &Config{
		Providers: make(map[string]Provider),
	}
}

func (c *Config) Load() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("getting home directory: %w", err)
	}

	configFile := filepath.Join(homeDir, ".aigit", "config.json")
	configData, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Return empty config if file doesn't exist
		}
		return fmt.Errorf("reading config file: %w", err)
	}

	if err := json.Unmarshal(configData, c); err != nil {
		return fmt.Errorf("parsing config file: %w", err)
	}

	return nil
}

func (c *Config) Save() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("getting home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".aigit")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return fmt.Errorf("creating config directory: %w", err)
	}

	configFile := filepath.Join(configDir, "config.json")
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("encoding config: %w", err)
	}

	if err := os.WriteFile(configFile, jsonData, 0600); err != nil {
		return fmt.Errorf("saving config: %w", err)
	}

	return nil
}

func (c *Config) AddProvider(provider, apiKey string, endpoint ...string) error {
	switch provider {
	case ProviderDoubao:
		c.Providers[provider] = Provider{
			APIKey:   apiKey,
			Endpoint: endpoint[0],
		}
	default:
		c.Providers[provider] = Provider{
			APIKey: apiKey,
		}
	}
	if c.CurrentProvider == "" {
		c.CurrentProvider = provider
	}
	return c.Save()
}

func (c *Config) UseProvider(provider string) error {
	if _, exists := c.Providers[provider]; !exists {
		return fmt.Errorf("provider %s not configured", provider)
	}
	c.CurrentProvider = provider
	return c.Save()
}

func (c *Config) GetAPIKey(provider string) (string, error) {
	if p, exists := c.Providers[provider]; exists {
		return p.APIKey, nil
	}
	return "", fmt.Errorf("no API key found for provider %s", provider)
}

func (c *Config) ListProviders() []string {
	providers := make([]string, 0, len(c.Providers))
	for k, _ := range c.Providers {
		if k == c.CurrentProvider {
			providers = append(providers, k+" *default")
		} else {
			providers = append(providers, k)
		}
	}
	return providers
}
