package main

import (
	"gitlab.com/Tobbeman/script-runner/internal/config"
	"gitlab.com/Tobbeman/script-runner/internal/handler"
	"gitlab.com/Tobbeman/script-runner/internal/runner"
	"gitlab.com/Tobbeman/script-runner/internal/server"
	"log"
)

func main(){
	conf, err := config.Setup(config.DefaultConfigPath)
	if err != nil {
		log.Fatal("error opening config: ", err)
	}
	log.Print("Loaded config, token is ", conf.Token)

	r := runner.New(conf.ScriptPath)
	h := handler.New(r, conf)

	http := server.New()
	root := http.Group("")

	h.Register(root)

	http.Start("0.0.0.0:8080")
}


