package zplugin

import (
	"plugin"

	"github.com/cocktail828/gdk/v1/logger"
	"github.com/cocktail828/gdk/v1/must"
	"github.com/cocktail828/gdk/v1/responsibe_chain"
)

var _pluginManager *responsibe_chain.ResponsibeChain

func init() {
	_pluginManager = responsibe_chain.New()
}

func InitPluginManager(confStr string, names ...string) {
	logger.Default().Println("pre init plugins: ", names)
	if len(names) == 0 {
		logger.Default().Fatal("empty plugins")
	}

	for _, v := range names {
		logger.Default().Println("load plugin: ", v)
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

		_pluginManager.Register(f.(func() ZPlugin)())
	}

	_pluginManager.Traverse(func(h responsibe_chain.Handler) {
		p := h.(ZPlugin)
		must.Must(p.Init(confStr))
		logger.Default().Println(p.Name(), "init success")
	})
}

func Plugins(f func(res responsibe_chain.Handler)) {
	_pluginManager.Traverse(f)
}
