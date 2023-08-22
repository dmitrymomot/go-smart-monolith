package main

import (
	"fmt"
	"net/http"

	"github.com/dmitrymomot/go-smart-monolith/internal/user/ports/restapi"
	"github.com/dmitrymomot/go-smart-monolith/internal/user/service"
	"github.com/dmitrymomot/go-smart-monolith/pkg/logx"
	"github.com/dmitrymomot/go-smart-monolith/pkg/nats"
	"github.com/dmitrymomot/go-smart-monolith/pkg/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// init router
	// Using chi router here, but you can use any other router as well.
	r := chi.NewRouter()
	// ...Set up all global middlewares your app needs here, like
	// request logging, tracing, auth, cors, etc.
	// Some more specific middlewares might need to be set on
	// the individual routes in the services transport layer.
	r.Use(middleware.Recoverer)

	// init logger
	// Using wrapper instead of direct logger initialization
	// to be able to change logger implementation in the future.
	log := logx.New()

	// init user app storage
	// low-level storage implementation, that can be used by service adapters,
	// like repositories, etc.
	stor := storage.New()

	// init nats client
	nats := nats.NewClient()

	// mount user service
	r.Mount("/users", restapi.NewServer(service.NewService(
		stor, log, nats,
		service.Config{
			// ...Set up all service-specific configs here.
		},
	)))

	// ...Mount more services here.

	// start server
	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
