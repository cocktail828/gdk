package main

import (
	"github.com/cocktail828/gdk/v1/errcode"
	"github.com/cocktail828/gdk/v1/message"
	"github.com/cocktail828/gdk/v1/zplugin"
)

var _ zplugin.ZPlugin = &verbose{}

type verbose struct{}

func New() zplugin.ZPlugin {
	return &verbose{}
}

func (v *verbose) Name() string {
	return "verbose"
}

func (v *verbose) Interest(msg *message.Message) bool {
	return true
}

func (v *verbose) Init(cfgs map[string][]byte) *errcode.Error {
	return nil
}

func (v *verbose) Preproc(msg *message.Message, tools zplugin.Tools) (zplugin.State, *errcode.Error) {
	// tools.Logger().Info()
	return zplugin.CONT, nil
}

func (v *verbose) Process(msg *message.Message, tools zplugin.Tools) (zplugin.State, *errcode.Error) {
	return zplugin.CONT, nil
}

func (v *verbose) Postproc(msg *message.Message, tools zplugin.Tools) (zplugin.State, *errcode.Error) {
	return zplugin.CONT, nil
}
