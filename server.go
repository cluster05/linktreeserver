package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/cluster05/linktree/api"
	"github.com/cluster05/linktree/api/config"
)

//	@title			linktree documentation API
//	@version		1.0.0
//	@description	linktree is platform of the new generation to extend their reach to new world with single link.

//	@host		localhost:3000
//	@BasePath	/

//	@accept						json
//	@produce					json
//
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	if err := setupEnvironment(); err != nil {
		return err
	}

	router, err := api.InitRouter()
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.ServerConfig.Port),
		Handler:        router,
		ReadTimeout:    config.ServerConfig.ReadTimeout * time.Second,
		WriteTimeout:   config.ServerConfig.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func setupEnvironment() error {

	env := flag.String("env", "dev", "To set environment dev/stg/prod")
	flag.Parse()

	var configFile string

	if *env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		configFile = "config.json"
	} else if *env == "stg" {
		configFile = "config_stg.json"
	} else {
		configFile = "config_dev.json"
	}

	if err := config.Setup(configFile); err != nil {
		return err
	}

	return nil
}
