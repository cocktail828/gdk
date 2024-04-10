package verbose

import (
	"github.com/cocktail828/gdk/v1/errcode"
	"github.com/cocktail828/gdk/v1/zplugin"
	"github.com/cocktail828/gdk/v1/zplugin/consts"
)

var _ zplugin.ZPlugin = &verbose{}

type verbose struct{}

func (v *verbose) Name() string {
	return "verbose"
}

func (v *verbose) Init(cfg []byte) *errcode.Error {
	return nil
}

func (v *verbose) Interest(ctx *zplugin.Context) bool {
	return true
}

func (v *verbose) Prepare(ctx *zplugin.Context) (consts.State, *errcode.Error) {
	// tools.Logger().Info()
	return consts.CONT, nil
}

func (v *verbose) Process(ctx *zplugin.Context) (consts.State, *errcode.Error) {
	return consts.CONT, nil
}

func (v *verbose) Cleanup(ctx *zplugin.Context) (consts.State, *errcode.Error) {
	return consts.CONT, nil
}
