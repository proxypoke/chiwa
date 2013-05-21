// chiwa - change i3 wallpapers automatically
//
// Author: slowpoke <mail+git@slowpoke.io>
//
// This program is free software under the non-terms
// of the Anti-License. Do whatever the fuck you want.
//
// Github: https://www.github.com/proxypoke/chiwa
// (Shortlink: https://git.io/chiwa)

package wallpaper

// TODO: Several of the tests assume that the primary display is LVDS1.

import (
	"testing"

	"github.com/BurntSushi/xgb/randr"

	//"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xgraphics"
)

func TestNewBG(t *testing.T) {
	X, err := xgbutil.NewConn()
	if err != nil {
		t.Errorf("%s\n", err)
	}
	randr.Init(X.Conn())
	_, err = NewBackground(X)
	if err != nil {
		t.Errorf("%s\n", err)
	}
}

func setup() *xgbutil.XUtil {
	X, _ := xgbutil.NewConn()
	randr.Init(X.Conn())
	return X
}


func TestSetImg(t *testing.T) {
	X := setup()
	bg, err := NewBackground(X)

	img, err := xgraphics.NewFileName(X, "kos-mos.png")
	if err != nil {
		t.Errorf("%s\n", err)
	}
	if err = SetImageToBg(X, bg, img, "LVDS1"); err != nil {
		t.Errorf("%s\n", err)
	}

	//win := bg.XShowExtra("test", true)
	//bg.XPaint(win.Id)
	//xevent.Main(X)
}

func TestSetRoot(t *testing.T) {
	X := setup()
	bg, err := NewBackground(X)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	// Errors in these functions are tested by TestSetImg, so we ignore them.
	img, _ := xgraphics.NewFileName(X, "kos-mos.png")
	SetImageToBg(X, bg, img, "LVDS1")

	if err = SetRoot(X, bg); err != nil {
		t.Errorf("%s\n", err)
	}
}

