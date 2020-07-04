package main

import (
	"flag"
	"gitlab.com/Tobbeman/script-runner/internal/config"
	"gitlab.com/Tobbeman/script-runner/internal/handler"
	"gitlab.com/Tobbeman/script-runner/internal/runner"
	"gitlab.com/Tobbeman/script-runner/internal/server"
	"log"
)

func main(){
	var configPath = flag.String("config-path", config.DefaultConfigPath, "filepath to config")
	flag.Parse()

	conf, err := config.Setup(*configPath)
	if err != nil {
		log.Fatal("error opening config: ", err)
	}
	log.Print("Loaded config, token is ", conf.Token)

	r := runner.New(conf.ScriptPath)
	h := handler.New(r, conf)

	http := server.New()
	root := http.Group("")

	h.Register(root)

	if err := http.Start(conf.Address); err != nil {
		log.Fatal(err)
	}
}


