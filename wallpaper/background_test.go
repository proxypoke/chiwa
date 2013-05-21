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

func TestSetImg(t *testing.T) {
	X, err := xgbutil.NewConn()
	if err != nil {
		t.Errorf("%s\n", err)
	}
	randr.Init(X.Conn())
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
