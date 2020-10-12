package client

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

const (
	ConfigKeyNamespace = "namespace"
	ConfigKeyFeatures  = "features"
	ConfigKeyCACert    = "cacert"
)

// KuberaConfig is a map of strings to interface{} for deserializing Kubera client config options.
// The alias is a way to attach type-asserting convenience methods.
type KuberaConfig map[string]interface{}

// LoadConfig loads the Kubera client configuration file and returns it as a KuberaConfig. If the
// file does not exist, an empty map is returned.
func LoadConfig() (KuberaConfig, error) {
	fileName := configFileName()

	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		// If the file isn't there, just return an empty map
		return KuberaConfig{}, nil
	}
	if err != nil {
		// For any other Stat() error, return it
		return nil, errors.WithStack(err)
	}

	configFile, err := os.Open(fileName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer configFile.Close()

	var config KuberaConfig
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		return nil, errors.WithStack(err)
	}

	return config, nil
}

// SaveConfig saves the passed in config map to the Kubera client configuration file.
func SaveConfig(config KuberaConfig) error {
	fileName := configFileName()

	// Try to make the directory in case it doesn't exist
	dir := filepath.Dir(fileName)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return errors.WithStack(err)
	}

	configFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return errors.WithStack(err)
	}
	defer configFile.Close()

	return json.NewEncoder(configFile).Encode(&config)
}

func (c KuberaConfig) Namespace() string {
	val, ok := c[ConfigKeyNamespace]
	if !ok {
		return ""
	}

	ns, ok := val.(string)
	if !ok {
		return ""
	}

	return ns
}

// Features return the feature gates provided
// by user via configurations
func (c KuberaConfig) Features() []string {
	val, ok := c[ConfigKeyFeatures]
	if !ok {
		return []string{}
	}

	features, ok := val.(string)
	if !ok {
		return []string{}
	}

	return strings.Split(features, ",")
}

func (c KuberaConfig) CACertFile() string {
	val, ok := c[ConfigKeyCACert]
	if !ok {
		return ""
	}
	caCertFile, ok := val.(string)
	if !ok {
		return ""
	}

	return caCertFile
}

func configFileName() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "kubera", "config.json")
}
