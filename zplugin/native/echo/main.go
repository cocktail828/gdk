package main

import (
	"github.com/cocktail828/gdk/v1/errcode"
	"github.com/cocktail828/gdk/v1/message"
	"github.com/cocktail828/gdk/v1/zplugin"
)

var _ zplugin.ZPlugin = &echo{}

type echo struct{}

func New() zplugin.ZPlugin {
	return &echo{}
}

func (d *echo) Name() string {
	return "echo"
}

func (d *echo) Interest(msg *message.Message) bool {
	return true
}

func (d *echo) Init(cfgs map[string][]byte) *errcode.Error {
	return nil
}

func (d *echo) Preproc(msg *message.Message, tools zplugin.Tools) (zplugin.State, *errcode.Error) {
	return zplugin.CONT, nil
}

func (d *echo) Process(msg *message.Message, tools zplugin.Tools) (zplugin.State, *errcode.Error) {
	return zplugin.CONT, nil
}

func (d *echo) Postproc(msg *message.Message, tools zplugin.Tools) (zplugin.State, *errcode.Error) {
	return zplugin.CONT, nil
}
