package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Web struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Timeout struct {
			Server time.Duration `yaml:"server"`
			Write  time.Duration `yaml:"write"`
			Read   time.Duration `yaml:"read"`
		} `yaml:"timeout"`
	} `yaml:"web"`

	Database struct {
		User       string `yaml:"user"`
		Password   string `yaml:"password"`
		Host       string `yaml:"host"`
		Name       string `yaml:"name"`
		DisableTLS bool   `yaml:"disabletls"`
	} `yaml:"database"`

	Redis struct {
		Address string `yaml:"address"`
		Port    string `yaml:"port"`
		DB      int    `yaml:"db"`
	} `yaml:"redis"`
}

// Config struct for webapp config

//flags struct for CLI configs change
//Might use it for future work
type Flags struct {
	Path    string
	Seed    string
	Migrate string
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() (Flags, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	if err := ValidateConfigPath(configPath); err != nil {
		return Flags{}, err
	}

	// Return the configuration path
	return Flags{
		Path: configPath,
	}, nil
}
