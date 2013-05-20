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

	"github.com/BurntSushi/xgb/randr"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/BurntSushi/xgbutil/xrect"
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
		oinfo, err := randr.GetOutputInfo(X.Conn(), output, 0).Reply()
		if err != nil {
			return nil, err
		}
		if oinfo.Crtc == 0 {
			// this output is disabled
			continue
		}
		crtc, err := randr.GetCrtcInfo(X.Conn(), oinfo.Crtc, 0).Reply()
		if err != nil {
			return nil, err
		}

		x, y := int(crtc.X), int(crtc.Y)
		w, h := int(crtc.Width), int(crtc.Height)
		r := image.Rectangle{
			image.Point{x, x + w},
			image.Point{y, y + h},
		}

		bgRect = bgRect.Union(r)
	}
	return xgraphics.New(X, bgRect), nil
}
