package server

import (
	"sync"
)

type Server interface {
	Name() string
	Init(conf string, opt ...string) error
	Run(wg *sync.WaitGroup) error
}
