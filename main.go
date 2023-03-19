package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/caarlos0/env"
	"k8s.io/klog"
)

type config struct {
	Port      int           `env:"PORT" envDefault:"8080"`
	JwtSecret string        `env:"JWT_SECRET" envDefault:"test"`
	JwtExpire time.Duration `env:"JWT_EXPIRE" envDefault:"24h"`
}

var cfg = config{}

func init() {
	flag.Parse()
	if err := env.Parse(&cfg); err != nil {
		log.Panic(err)
	}
}

func main() {
	runEcho(cfg.Port)
	klog.Info("fire")
	ctlc := make(chan os.Signal, 1)
	signal.Notify(ctlc, os.Interrupt)
	<-ctlc
	klog.Info("done")
}
