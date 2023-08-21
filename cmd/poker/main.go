package main

import (
	"context"
	"github.com/devandreyl/go-poker-hands-evaluator/cmd/poker/handler"
	"github.com/go-playground/validator/v10"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func main() {
	var (
		ctx, stop = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		config    = MustReadConfig()
	)

	router := mux.NewRouter()
	evaluateHandler := handler.NewEvaluateHandler(router, validator.New())
	evaluateHandler.Register()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	routerWithCORS := corsMiddleware.Handler(router)

	server := CreateHTTPServer(config, routerWithCORS)

	run(ctx, stop, server)
}

func run(
	ctx context.Context,
	stop context.CancelFunc,
	server *http.Server,
) {
	errCh := make(chan error)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			select {
			case <-ctx.Done():
			case errCh <- errors.Wrap(err, "http server"):
			}
		}
	}()

	shutdown := func(err error) {
		_, cancelTimeout := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancelTimeout()

		stop()

		if err != nil {
			log.Panic("shutdown caused by error")
		}
	}

	select {
	case <-ctx.Done():
		shutdown(nil)
	case err := <-errCh:
		shutdown(err)
	}
}
