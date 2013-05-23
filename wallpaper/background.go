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
	"image/draw"

	"github.com/proxypoke/chiwa/util"

	"github.com/BurntSushi/xgb/randr"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/BurntSushi/xgbutil/xprop"
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
		// NOTE: It doesn't really matter if this returns a Zero Rectangle.
		if err != nil {
			return nil, err
		}
		bgRect = bgRect.Union(r)
	}
	return xgraphics.New(X, bgRect), nil
}

// SetImageToBg sets the given image into the background at the proper location
// for the named output to display.
func SetImageToBg(X *xgbutil.XUtil,
	bg *xgraphics.Image,
	img image.Image,
	name string) error {

	output, err := util.GetOutputByName(X, name)
	if err != nil {
		return err
	}
	geom, err := util.OutputRect(X, output)
	if err != nil {
		return err
	}
	if err = bg.CreatePixmap(); err != nil {
		return err
	}
	bg.XDraw()
	draw.Draw(bg, geom, img, img.Bounds().Min, draw.Src)
	return nil
}

// SetRoot sets the given background as the root window background.
func SetRoot(X *xgbutil.XUtil, bg *xgraphics.Image) error {
	root := X.RootWin()
	if err := bg.XSurfaceSet(root); err != nil {
		return err
	}
	bg.XDraw()
	bg.XPaint(root)
	// FIXME: This doesn't set the pixmap persistently. As soon as the program
	// exits, the pixmap is destroyed. Find a way to make it persistent.
	xprop.ChangeProp32(X, root, "_XROOTPMAP_ID", "PIXMAP", uint(bg.Pixmap))
	xprop.ChangeProp32(X, root, "ESETROOT_PMAP_ID", "PIXMAP", uint(bg.Pixmap))
	return nil
}
