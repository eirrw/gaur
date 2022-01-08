package cfg

import (
	_ "embed"
	"errors"
	"github.com/BurntSushi/toml"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
)

const configFile = "gaur/config.toml"

//go:embed "config.toml.tmpl"
var configTemplateSource string
var configTemplate = template.Must(template.New("config").Parse(configTemplateSource))

// config contains the configuration information of the application
type config struct {
	CacheDir *string
	Remote	 *string
}

// New will generate a config struct with default values
func New() (*config, error) {
	cfg := config{}

	cache, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	cacheDir := filepath.Join(cache, "aur")
	remote := "origin"

	cfg.CacheDir = &cacheDir
	cfg.Remote = &remote

	return &cfg, nil
}

// Initialize will create the default config file
func Initialize(c *config) error {
	def, err := New()
	if err != nil {
		return err
	}

	// merge passed config with default values
	if c.CacheDir == nil {
		c.CacheDir = def.CacheDir
	}

	err = writeConfig(c)
	if err != nil {
		return err
	}

	return nil
}

// GetConfig reads the current config from file and returns it as a struct
func GetConfig() (*config, error) {
	var cfg config

	path, err := getConfigFilepath()
	if err != nil {
		return nil, err
	}

	if is, err := IsInit(); err != nil {
		return nil, err
	} else if !is {
		return nil, errors.New("config is not initialized, run \"gaur init\" to resolve")
	}

	_, err = toml.DecodeFile(path, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// writeConfig writes the config struct to the filesystem
func writeConfig(c *config) error {
	path, err := getConfigFilepath()
	if err != nil {
		return err
	}

	if err = os.MkdirAll(filepath.Dir(path), 0744); err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := configTemplate.Execute(file, c); err != nil {
		return err
	}

	return nil
}

// getConfigFilepath returns the path of the config file
func getConfigFilepath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	configPath := filepath.Join(configDir, configFile)

	return configPath, nil
}

// IsInit checks if the config file has been initialized
func IsInit() (bool, error) {
	path, err := getConfigFilepath()
	if err != nil {
		return false, err
	}

	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
