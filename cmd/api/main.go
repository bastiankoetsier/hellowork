package main

import (
	"fmt"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/italolelis/hellowork/config"
	"github.com/italolelis/hellowork/middleware"
	"github.com/italolelis/hellowork/router"
)

var (
	err          error
	globalConfig *config.Specification
	accessor     *middleware.DatabaseAccessor
)

// initializes the global configuration
func init() {
	globalConfig, err = config.LoadEnv()
	if nil != err {
		log.Panic(err.Error())
	}
}

// initializes the basic configuration for the log wrapper
func init() {
	level, err := log.ParseLevel(strings.ToLower(globalConfig.LogLevel))
	if err != nil {
		log.Error("Error getting level", err)
	}

	log.SetLevel(level)
}

// initializes a DB connection
func init() {
	accessor, err = middleware.InitDB(globalConfig.Database.DSN)
	if err != nil {
		log.Fatalf("Couldn't connect to the mongodb database: %s", err.Error())
	}
}

func main() {
	defer accessor.Close()

	router := router.NewHttpTreeMuxRouter()
	router.Use(middleware.NewLogger(globalConfig.Debug).Handler, middleware.NewRecovery(RecoveryHandler).Handler, middleware.NewMongoDB(accessor).Handler)

	// Home endpoint for the gateway
	router.GET("/", Home(globalConfig.Application))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", globalConfig.Port), router))
}
