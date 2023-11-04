package task

import (
	"bytes"
	"sync"

	"github.com/cocktail828/go-tools/z/reflectx"
)

var (
	defaultBufSize = 1024 * 100
	pool           = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, defaultBufSize))
		},
	}
)

func get() *bytes.Buffer {
	return pool.Get().(*bytes.Buffer)
}

func put(v *bytes.Buffer) {
	if !reflectx.IsNil(v) {
		pool.Put(v)
	}
}
