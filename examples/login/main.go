package main

import (
	"flag"
	log "github.com/thinkboy/log4go"
	"runtime"
)

var (
	Debug bool
	Ver   = "0.1"
)


func main() {
	flag.Parse()
	if err := InitConfig(); err != nil {
		panic(err)
	}
	Debug = Conf.Debug
	runtime.GOMAXPROCS(Conf.MaxProc)
	log.LoadConfiguration(Conf.Log)

	defer log.Close()

	log.Info("login %s start",Ver)
}





