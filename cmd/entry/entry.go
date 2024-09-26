package entry

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/cocktail828/gdk/v1/cmd/status"
	"github.com/cocktail828/gdk/v1/pkg/loader"
	"github.com/cocktail828/gdk/v1/servers"
	"github.com/cocktail828/go-tools/configor"
	"github.com/cocktail828/go-tools/z"
	"github.com/cocktail828/go-tools/z/reflectx"
	"golang.org/x/sync/errgroup"

	_ "github.com/cocktail828/gdk/v1/servers/restapi"
)

type Config struct {
	Servers struct {
		Servers []string `toml:"servers"`
	} `toml:"servers"`
	LogConfig struct {
		LogLevel  string `toml:"level" required:"true"`
		LogFile   string `toml:"file" required:"true"`
		LogSize   int    `toml:"size" default:"100"`
		LogCount  int    `toml:"count" default:"30"`
		LogCaller bool   `toml:"caller" default:"false"`
		LogAsync  bool   `toml:"async" default:"true"`
		LogBatch  int    `toml:"batch"`
		IDC       string `toml:"idc"`
		Bcluster  string `toml:"bcluster"`
	} `toml:"log" required:"true"`
}

func Start(inMode int, inCfg, inGroup, inService, inRegistry string) {
	{
		status.InMode = inMode
		status.InCfg = inCfg
		status.InGroup = inGroup
		status.InService = inService
		status.InRegistry = inRegistry
	}

	log.Printf("%s->inMode:%v,Group:%v,Service:%v,Version:%v,API:%v,Registry:%v",
		status.SrvName, inMode, inGroup, inService, status.AppVersion, status.ApiVersion, inRegistry)

	cfgor, err := loader.NewLoader(inMode, inCfg, inGroup, inService, inRegistry)
	z.Must(err)

	cfg := Config{}
	z.Must(configor.Load(&cfg, cfgor.GetRawCfg()))
	srvs := []servers.Server{}
	for _, srvname := range cfg.Servers.Servers {
		if val := servers.Lookup(srvname); !reflectx.IsNil(val) {
			srvs = append(srvs, val)
		} else {
			log.Printf("server %v is not implement", srvname)
		}
	}

	log.Println("about to starting metrics.")

	log.Println("about to starting servers.")
	sigctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	egrp, ctx := errgroup.WithContext(sigctx)
	egrp.SetLimit(len(cfg.Servers.Servers))
	for _, srv := range srvs {
		if err := srv.Init(cfgor.GetRawCfg(), nil); err != nil {
			log.Fatalf("init server:%v fail for:%v", srv.Name(), err)
		}
	}

	for _, srv := range srvs {
		_s := srv
		egrp.Go(func() error { return _s.Run(ctx) })
	}

	egrp.Wait()
}
