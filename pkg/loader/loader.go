package loader

import "sync"

//go:generate stringer -type Mode -linecomment
type Mode int

const (
	Native Mode = iota
	Centre Mode = iota
)

type configloader struct {
	mu  sync.RWMutex
	raw []byte
}

func NewLoader(inMode int, inCfg, inGroup, inService, inRegistry string) (*configloader, error) {
	return nil, nil
}

func (l *configloader) GetRawCfg() []byte {
	return l.raw
}
