package zplugin

import (
	"container/list"
	"log"
	"plugin"
	"strings"
	"sync/atomic"

	"github.com/cocktail828/go-tools/z"
	"github.com/cocktail828/go-tools/z/chain"
)

var (
	pluginManager *chain.Chain = chain.New()
)

func InitPlugins(cfgs map[string][]byte) {
	pluginManager.Traverse(nil, func(h chain.Handler) bool {
		w := h.(*wrapperedHandler)
		if w.inited.CompareAndSwap(false, true) {
			z.Must(w.Init(cfgs))
			log.Println("init buildin plugins success", w.Name())
		}
		return true
	})
}

type wrapperedHandler struct {
	ZPlugin
	inited   atomic.Bool
	isNative bool
}

func RegisterBuildinPlugins(plugins ...ZPlugin) {
	log.Println("try init buildin plugins", func() string {
		tmp := []string{}
		for _, s := range plugins {
			tmp = append(tmp, s.Name())
		}
		return strings.Join(tmp, ",")
	}())

	for _, p := range plugins {
		if p == nil || pluginManager.Get(p.Name()) != nil {
			continue
		}
		log.Println("load buildin plugin", p.Name())
		pluginManager.Add(&wrapperedHandler{
			ZPlugin:  p,
			isNative: false,
		})
	}
}

func RegisterNativePlugins(cfgs map[string][]byte, names ...string) {
	log.Println("try init native plugins", strings.Join(names, ","))
	for _, n := range names {
		if n == "" || pluginManager.Get(n) != nil {
			continue
		}

		log.Println("load native plugin", n)
		p, err := plugin.Open(n)
		if err != nil {
			log.Fatal(err)
		}

		f, err := p.Lookup(New)
		if err != nil {
			log.Fatal(err)
		}
		pluginManager.Add(&wrapperedHandler{
			ZPlugin:  f.(func() ZPlugin)(),
			isNative: false,
		})
	}
}

func Traverse(f func(p ZPlugin)) *list.Element {
	return pluginManager.Traverse(nil, func(h chain.Handler) bool {
		f(h.(*wrapperedHandler).ZPlugin)
		return true
	})
}

func Reverse(f func(p ZPlugin)) *list.Element {
	return pluginManager.Reverse(nil, func(h chain.Handler) bool {
		f(h.(*wrapperedHandler).ZPlugin)
		return true
	})
}
