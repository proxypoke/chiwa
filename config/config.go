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
	"strings"
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
