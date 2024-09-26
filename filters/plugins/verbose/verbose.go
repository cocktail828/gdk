package main

import (
	"github.com/cocktail828/gdk/v1/errcode"
	"github.com/cocktail828/gdk/v1/filters"
)

var _ filters.Filter = &verbose{}

type verbose struct{}

func NewFilter() filters.Filter {
	return &verbose{}
}

func (v *verbose) Name() string {
	return "verbose"
}

func (v *verbose) Init(data []byte) *errcode.Error {
	return nil
}

func (v *verbose) Interest(ctx *filters.Context) bool {
	return true
}

func (v *verbose) Prepare(ctx *filters.Context) *errcode.Error {
	return nil
}

func (v *verbose) Process(ctx *filters.Context) *errcode.Error {
	return nil
}

func (v *verbose) Cleanup(ctx *filters.Context) *errcode.Error {
	return nil
}
