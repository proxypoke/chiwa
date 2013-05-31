// chiwa - change i3 wallpapers automatically
//
// Author: slowpoke <mail+git@slowpoke.io>
//
// This program is free software under the non-terms
// of the Anti-License. Do whatever the fuck you want.
//
// Github: https://www.github.com/proxypoke/chiwa
// (Shortlink: https://git.io/chiwa)

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/proxypoke/chiwa/logging"
)

var (
	configName string = "chiwarc.json"
)

// Read tries reading a config from the given file.
func Read(path string) (Config, error) {
	var c Config
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(content, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}

// Write takes a config and writes it to the file.
func Write(path string, conf *Config) error {
	toWrite, err := json.Marshal(conf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, toWrite, os.FileMode(0744))
}

// LocateConfig tries to find a configuration file for chiwa, returning its
// filepath or an empty string.
func LocateConfig() string {
	var configPath string
	// First, look in the user's home directory.
	home := os.Getenv("HOME")
	logging.Debugf("$HOME == %s\n", home)
	if home == "" {
		logging.Warnf("$HOME is unset (this is not good).")
		return ""
	} else {
		// Try the standard location first.
		configPath = path.Join(home, "." + configName)
		logging.Debugf("Trying %s...\n", configPath)
		if fileExists(configPath) {
			return configPath
		}

		// Try $HOME/.config/chiwa.
		configPath = path.Join(home, ".config", "chiwa", configName)
		if fileExists(configPath) {
			return configPath
		}
	}

	// Finally, look in /etc.
	configPath = path.Join("/etc", configName)
	if fileExists(configPath) {
		return configPath
	}
	// We haven't found anything.
	return ""
}

// Check if a file exists.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

// Verify checks the given config for invalid values.
func Verify(c Config) error {
	for _, output := range c.Wallpaper {
		err := verifyOutput(output)
		if err != nil {
			return err
		}
	}
	return nil
}

// Helper function to verify a single output.
func verifyOutput(o Output) error {
	if !valid[o.Mode] {
		return fmt.Errorf("Invalid mode: %s (Expected one of %s)",
			o.Mode,
			func() string {
				var s []string
				for key := range valid {
					s = append(s, "'"+string(key)+"'")
				}
				return strings.Join(s, ", ")
			}())
	}
	return nil
}
