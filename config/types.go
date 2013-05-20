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

// BGMode represents a way in which a background image can be set.
// TODO: Move this to a more appropriate location. It doesn't belong in the
// config package.
type BGMode string

const (
	BGCenter BGMode = "center"
	BGFill          = "fill"
	BGMax           = "max"
	BGScale         = "scale"
	BGTile          = "tile"
)

// A set for internal verification.
var valid = map[BGMode]bool{
	BGCenter: true,
	BGFill:   true,
	BGMax:    true,
	BGScale:  true,
	BGTile:   true,
}

// Config represents a chiwa configuration.
type Config struct {
	// There are no global settings yet, but we include it here to be more
	// flexible later.
	Global    map[string]interface{}
	Wallpaper map[string]Output
}

// Output represents a single output as named by xrandr (e.g. VGA1 or LVDS1).
type Output struct {
	Mode   BGMode
	Image  string
	Images map[string]string
}
