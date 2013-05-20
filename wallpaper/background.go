// chiwa - change i3 wallpapers automatically
//
// Author: slowpoke <mail+git@slowpoke.io>
//
// This program is free software under the non-terms
// of the Anti-License. Do whatever the fuck you want.

// Github: https://www.github.com/proxypoke/chiwa
// (Shortlink: https://git.io/chiwa)

package wallpaper

import (
	"image"

	"github.com/proxypoke/chiwa/util"

	"github.com/BurntSushi/xgb/randr"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xgraphics"
)

// NewBackground creates an xgraphics.Image which spans the entire screen,
// initialized to black.
func NewBackground(X *xgbutil.XUtil) (*xgraphics.Image, error) {
	res, err := randr.GetScreenResources(X.Conn(), X.RootWin()).Reply()
	if err != nil {
		return nil, err
	}
	var bgRect image.Rectangle
	for _, output := range res.Outputs {
		r, err := util.OutputRect(X, output)
		if err != nil {
			return nil, err
		}
		bgRect = bgRect.Union(r)
	}
	return xgraphics.New(X, bgRect), nil
}
