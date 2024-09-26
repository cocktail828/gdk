package filters

import (
	"plugin"
	"strings"
	"sync"

	"github.com/cocktail828/gdk/v1/errcode"
	"github.com/cocktail828/go-tools/log"
	"github.com/cocktail828/go-tools/z"
	"github.com/pkg/errors"
)

type Filter interface {
	Name() string
	Init(data []byte) *errcode.Error
	Interest(ctx *Context) bool
	Prepare(ctx *Context) *errcode.Error
	Process(ctx *Context) *errcode.Error
	Cleanup(ctx *Context) *errcode.Error
}

var (
	pluginMu  sync.RWMutex
	pluginMap = map[string]Filter{}
)

func lookupPlugin(name string) Filter {
	pluginMu.RLock()
	defer pluginMu.RUnlock()
	if !strings.HasSuffix(name, ".so") {
		name += ".so"
	}
	return pluginMap[name]
}

// filename with suffix ".so"
func loadPlugin(name string) Filter {
	if name == "" {
		return nil
	}

	if !strings.HasSuffix(name, ".so") {
		name += ".so"
	}

	pluginMu.Lock()
	defer pluginMu.Unlock()
	if p, ok := pluginMap[name]; ok {
		return p
	}

	log.Println(">>> try load plugin:", name)
	p, err := plugin.Open(name)
	z.Must(err)

	f, err := p.Lookup("NewFilter")
	z.Must(errors.WithMessagef(err, "plugin:%v, fail to lookup symbol:NewFilter", name))

	val := f.(func() Filter)()
	pluginMap[name] = val
	return val
}

type filterIMPL struct {
	PluginName string
	Filter     Filter
}

var (
	filterMu  = sync.RWMutex{}
	filterMap = map[string][]filterIMPL{}
)

func RegisterFilters(service string, pluginNames []string) {
	if len(service) == 0 {
		log.Fatalf("illegal service:%v", service)
	}

	if len(pluginNames) == 0 {
		return
	}

	log.Printf("about to register filters, service:%v,plugin:%v", service, pluginNames)
	filterMu.Lock()
	defer filterMu.Unlock()
	for _, name := range pluginNames {
		for _, f := range filterMap[service] {
			if f.PluginName == name {
				log.Fatalf("duplicate filter found, service:%v,plugin:%v", service, name)
			}
		}

		filterMap[service] = append(filterMap[service], filterIMPL{
			PluginName: name,
			Filter:     loadPlugin(name),
		})
	}

	log.Printf("register filters, service:%v,plugin:%v", service, pluginNames)
}

func GetFilters(service string, extras ...string) ([]Filter, *errcode.Error) {
	if len(service) == 0 {
		log.Printf("illegal service:%v", service)
		return nil, nil
	}

	filterMu.RLock()
	defer filterMu.RUnlock()
	chain := make([]Filter, 0, len(filterMap))
	if len(extras) != 0 {
		for _, pluginName := range extras {
			filter := lookupPlugin(pluginName)
			if filter == nil {
				return chain, errcode.CodeGeneralErr.WithMessage(pluginName + " is not supported")
			}
			chain = append(chain, filter)
		}
		return chain, nil
	}

	for _, f := range filterMap[service] {
		chain = append(chain, f.Filter)
	}
	return chain, nil
}
