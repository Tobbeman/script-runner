package main

import (
	"flag"
	"gitlab.com/Tobbeman/script-runner/internal/config"
	"gitlab.com/Tobbeman/script-runner/internal/handler"
	"gitlab.com/Tobbeman/script-runner/internal/runner"
	"gitlab.com/Tobbeman/script-runner/internal/server"
	"log"
)

func main() {
	var configPath = flag.String("config-path", config.DefaultConfigPath, "filepath to config")
	flag.Parse()

	conf, err := config.Setup(*configPath)
	if err != nil {
		log.Fatal("error opening config: ", err)
	}
	log.Print("Loaded config, token is ", conf.Token)

	r := runner.New(conf.ScriptPath)

	var s *runner.RCmdStore
	if conf.HasRetention() {
		s = runner.NewStoreWithRetention(conf.Retention)
	} else {
		s = runner.NewStore()
	}

	h := handler.New(r, s, conf)

	http := server.New()
	root := http.Group("")

	h.Register(root)

	if err := http.Start(conf.Address); err != nil {
		log.Fatal(err)
	}
}
