package main

import (
	"github.com/jackiiilong/leafserver/cluster"
	"github.com/jackiiilong/leafserver/componet/game"
	"github.com/jackiiilong/leafserver/componet/gate"
	"github.com/jackiiilong/leafserver/componet/login"
	"github.com/jackiiilong/leafserver/conf"
	"github.com/jackiiilong/leafserver/console"
	"github.com/jackiiilong/leafserver/log"
	"github.com/jackiiilong/leafserver/module"
	"os"
	"os/signal"
)

const version = "1.1.3"

func Run(mods ...module.IModule) {
	// logger
	if conf.LogLevel != "" {
		logger, err := log.New(conf.LogLevel, conf.LogPath, conf.LogFlag)
		if err != nil {
			panic(err)
		}
		log.Export(logger)
		defer logger.Close()
	}

	log.Release("Leaf %v starting up", version)

	// module
	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}


	module.Init()

	// cluster
	cluster.Init()

	// console
	console.Init()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Release("Leaf closing down (signal: %v)", sig)
	console.Destroy()
	cluster.Destroy()
	module.Destroy()
}

func main() {
	conf.LogLevel = conf.Server.LogLevel
	conf.LogPath = conf.Server.LogPath
	conf.LogFlag = conf.LogFlag
	conf.ConsolePort = conf.Server.ConsolePort
	conf.ProfilePath = conf.Server.ProfilePath

	Run(
		game.Game,
		gate.Gate,
		login.Login,
	)
}
