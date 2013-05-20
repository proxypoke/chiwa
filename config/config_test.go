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
	"io/ioutil"
	"os"
	"testing"
)

var (
	TEST_CONFIG string = "./test_config.json"
	EXPECTED           = &Config{
		map[string]interface{}{},
		map[string]Output{
			"LVDS1": Output{
				BGFill,
				"/path/to/image.png",
				map[string]string{},
			},
			"VGA1": Output{
				BGCenter,
				"",
				map[string]string{
					"1920x1080": "/path/to/image.png",
					"1280x720":  "/path/to/different/image.png",
					"default":   "/path/to/default/image.png",
				},
			},
		},
	}
)

func TestRead(t *testing.T) {
	_, err := Read(TEST_CONFIG)
	if err != nil {
		t.Errorf("Can't read config: %v", err)
	}
	// TODO: Actually check if the read config matches the expected config.
}

func TestWrite(t *testing.T) {
	tmp, err := ioutil.TempFile("", "chiwa_test_config")
	if err != nil {
		t.Errorf("Can't create a temporary file for testing: %v", err)
	}
	defer os.Remove(tmp.Name())
	err = Write(tmp.Name(), EXPECTED)
	if err != nil {
		t.Errorf("Can't write config: %v", err)
	}
}
