package responsibe_chain

import (
	"fmt"
	"strings"
	"sync"
)

type Handler interface {
	Name() string
}

type ResponsibeChain struct {
	mu         *sync.RWMutex
	handlerMap map[string]struct{}
	handlers   []Handler
}

func New() *ResponsibeChain {
	return &ResponsibeChain{
		mu:         &sync.RWMutex{},
		handlerMap: make(map[string]struct{}),
	}
}

func (rc ResponsibeChain) String() string {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	return fmt.Sprintf("handlers: [%v]", func() string {
		strs := []string{}
		for k := range rc.handlerMap {
			strs = append(strs, k)
		}
		return strings.Join(strs, ",")
	}())
}

func (rc *ResponsibeChain) Register(h Handler) *ResponsibeChain {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	if _, ok := rc.handlerMap[h.Name()]; ok {
		rc.handlerMap[h.Name()] = struct{}{}
		rc.handlers = append(rc.handlers, h)
	}
	return rc
}

func (rc *ResponsibeChain) Length() int {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	return len(rc.handlers)
}

func (rc *ResponsibeChain) Traverse(f func(h Handler)) {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	if f == nil {
		return
	}
	for _, h := range rc.handlers {
		f(h)
	}
}
