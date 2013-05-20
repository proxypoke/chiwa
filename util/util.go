// chiwa - change i3 wallpapers automatically
//
// Author: slowpoke <mail+git@slowpoke.io>
//
// This program is free software under the non-terms
// of the Anti-License. Do whatever the fuck you want.
//
// Github: https://www.github.com/proxypoke/chiwa
// (Shortlink: https://git.io/chiwa)

package util

import (
	"image"

	"github.com/BurntSushi/xgbutil"

	"github.com/BurntSushi/xgb/randr"
	//"github.com/BurntSushi/xgb/xproto"
)

func GetOutputs(X *xgbutil.XUtil) (outputs []*randr.GetOutputInfoReply) {
	resources, err := randr.GetScreenResources(X.Conn(), X.RootWin()).Reply()
	if err != nil {
		return nil
	}
	for _, output := range resources.Outputs {
		oinfo, err := randr.GetOutputInfo(X.Conn(), output, 0).Reply()
		if err != nil {
			return nil
		}
		outputs = append(outputs, oinfo)
	}
	return outputs
}

// OutputRect creates a Rectangle fitting the active CRTC of the given output.
// Returns the Zero Rectangle and an error if anything goes wrong. If the output
// is disabled, the ZR will be returned.
func OutputRect(X *xgbutil.XUtil, output randr.Output) (image.Rectangle, error) {
	oinfo, err := randr.GetOutputInfo(X.Conn(), output, 0).Reply()
	if err != nil {
		return image.ZR, err
	}
	if oinfo.Crtc == 0 {
		// this output is disabled
		return image.ZR, nil
	}
	crtc, err := randr.GetCrtcInfo(X.Conn(), oinfo.Crtc, 0).Reply()
	if err != nil {
		return image.ZR, err
	}

	x, y := int(crtc.X), int(crtc.Y)
	w, h := int(crtc.Width), int(crtc.Height)
	return image.Rectangle{
		image.Point{x, x + w},
		image.Point{y, y + h},
	}, nil
}
