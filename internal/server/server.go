package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/treddy73/go-fullstack/internal/server/route"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	*http.Server
}

func New(c Config) (Server, error) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Mount("/", route.Routes())

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static", fs))

	s := http.Server{
		Addr:    fmt.Sprintf(":%d", c.Port()),
		Handler: r,
	}

	return Server{&s}, nil
}

func (s Server) Start() {
	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	notifyCtx, notifyCtxStop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer notifyCtxStop()

	go func() {
		<-notifyCtx.Done()
		fmt.Println("server shutdown ...")

		shutdownCtx, shutdownCtxStop := context.WithTimeout(serverCtx, 5*time.Second)
		defer shutdownCtxStop()

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				fmt.Println("graceful shutdown timed out ... forcing exit ...")
				os.Exit(1)
			}
		}()

		if err := s.Shutdown(shutdownCtx); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		serverStopCtx()
	}()

	fmt.Printf("server listening on %s\n", s.Addr)

	if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		fmt.Println(err)
		os.Exit(1)
	}

	<-serverCtx.Done()
}
