package internal

import (
	"flag"
	"github.com/spf13/viper"
	"github.com/tidwall/buntdb"
	"time"
)

type Option struct {
	Host string
	Port string
	Ttl  int
}

func ParseOption() {
	var opt = new(Option)
	// for current set default
	viper.SetDefault("app.db", "memory")

	flag.StringVar(&opt.Host, "host", "localhost", "Running host server")
	flag.StringVar(&opt.Host, "h", "localhost", "Running host server")

	flag.StringVar(&opt.Port, "port", "9090", "Running port server")
	flag.StringVar(&opt.Port, "p", "9090", "Running port server")

	flag.IntVar(&opt.Ttl, "ttl", 1440, "How long data will keep in memory")
	flag.IntVar(&opt.Ttl, "t", 1440, "How long data will keep in memory")
	flag.Parse()

	viper.Set("app.host", opt.Host)
	viper.Set("app.port", opt.Port)

	var opts = new(buntdb.SetOptions)
	opts.Expires = false

	if opt.Ttl > 0 {
		opts.Expires = true
		opts.TTL = time.Duration(opt.Ttl) * time.Minute
	}
	viper.SetDefault("app.opt", opts)
}
