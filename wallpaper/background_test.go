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
	"github.com/BurntSushi/xgbutil"
)

func TestNewBG(t *testing.T) {
	X, err := xgbutil.NewConn()
	randr.Init(X.Conn())
	if err != nil {
		t.Errorf("%s\n", err)
	}
	_, err = NewBackground(X)
	if err != nil {
		t.Errorf("%s\n", err)
	}
}
