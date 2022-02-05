package main

import (
	"github.com/bmizerany/pat"
	"github.com/dragun-igor/chat-ws/internal/handlers"
	"net/http"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandleFunc(handlers.Home))

	return mux
}
