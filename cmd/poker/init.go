package main

import (
	"fmt"
	"github.com/devandreyl/go-poker-hands-evaluator/internal/config"
	"net/http"
)

func MustReadConfig() *config.Config {
	cnf, err := config.ReadConfig(true)
	if err != nil {
		panic(fmt.Sprintf("read config: %s", err))
	}

	return cnf
}

func CreateHTTPServer(
	cnf *config.Config,
	h http.Handler,
) *http.Server {
	return &http.Server{
		Handler: h,
		Addr:    cnf.HTTP.Listen,
	}
}
