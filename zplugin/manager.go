package zplugin

import (
	"container/list"
	"log"
	"plugin"

	"github.com/cocktail828/go-tools/z"
	"github.com/cocktail828/go-tools/z/chain"
)

var (
	pluginManager *chain.Chain = chain.New()
)

func InitPluginManager(confStr string, names ...string) {
	log.Println("pre init plugins: ", names)
	if len(names) == 0 {
		log.Fatal("empty plugins")
	}

	for _, v := range names {
		log.Println("load plugin: ", v)
		if v == "" {
			continue
		}

		p, err := plugin.Open(v)
		if err != nil {
			panic(err)
		}

		f, err := p.Lookup(New)
		if err != nil {
			panic(err)
		}
		pluginManager.Add(f.(func() ZPlugin)())
	}

	pluginManager.Traverse(nil, func(h chain.Handler) bool {
		p := h.(ZPlugin)
		z.Must(p.Init(confStr))
		log.Println(p.Name(), "init success")
		return true
	})
}

func Traverse(f func(res chain.Handler)) *list.Element {
	return pluginManager.Traverse(nil, func(h chain.Handler) bool {
		f(h)
		return true
	})
}

func Reverse(f func(res chain.Handler)) *list.Element {
	return pluginManager.Reverse(nil, func(h chain.Handler) bool {
		f(h)
		return true
	})
}

func Handlers() []chain.Handler {
	return pluginManager.Handlers()
}
