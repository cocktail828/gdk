package zplugin

import (
	"container/list"
	"log"
	"plugin"
	"strings"
	"sync/atomic"

	"github.com/cocktail828/gdk/v1/message"
	"github.com/cocktail828/go-tools/z"
	"github.com/cocktail828/go-tools/z/chain"
)

var (
	pluginManager *chain.Chain = chain.New()
)

func InitPlugins(cfgs map[string][]byte) {
	pluginManager.Traverse(nil, func(h chain.Handler) bool {
		w := h.(*wrappered)
		if w.inited.CompareAndSwap(false, true) {
			z.Must(w.Init(cfgs))
			log.Println("init buildin plugins success", w.Name())
		}
		return true
	})
}

type wrappered struct {
	ZPlugin
	inited   atomic.Bool
	isNative bool
}

func registerPlugin(w *wrappered) {
	pluginManager.Add(w)
	if v, ok := w.ZPlugin.(MessageParser); ok {
		message.RegisterParser(v.Sub(), v.Parser)
	}
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
		registerPlugin(&wrappered{
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
		registerPlugin(&wrappered{
			ZPlugin:  f.(func() ZPlugin)(),
			isNative: true,
		})
	}
}

func Traverse(f func(p ZPlugin)) *list.Element {
	return pluginManager.Traverse(nil, func(h chain.Handler) bool {
		f(h.(*wrappered).ZPlugin)
		return true
	})
}

func Reverse(f func(p ZPlugin)) *list.Element {
	return pluginManager.Reverse(nil, func(h chain.Handler) bool {
		f(h.(*wrappered).ZPlugin)
		return true
	})
}
