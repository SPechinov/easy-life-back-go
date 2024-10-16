package server

import (
	"easy-life-back-go/pkg/http_server"
	"fmt"
	"log/slog"
	"net/http"
)

func Start(port string) {
	s := http_server.New(port, "/api")

	s.Use(func(w http.ResponseWriter, r *http.Request, next http_server.Next) {
		fmt.Println("MW 1")
		next()
	})

	s.AddRoute(http.MethodPost, "/11", func(w http.ResponseWriter, r *http.Request) {})

	gQWE := s.NewGroup("/qwe")
	gQWEASD := gQWE.NewGroup("/asd")

	gQWE.Use(func(w http.ResponseWriter, r *http.Request, next http_server.Next) {
		fmt.Println("MW gQWE")
		next()
	})

	gQWEASD.Use(func(w http.ResponseWriter, r *http.Request, next http_server.Next) {
		fmt.Println("MW gQWEASD")
		next()
	})

	gQWE.AddRoute(http.MethodPost, "/11", func(w http.ResponseWriter, r *http.Request) {})
	gQWE.AddRoute(http.MethodGet, "/11", func(w http.ResponseWriter, r *http.Request) {})
	gQWE.AddRoute(http.MethodPatch, "/11", func(w http.ResponseWriter, r *http.Request) {})
	gQWE.AddRoute(http.MethodPost, "/22", func(w http.ResponseWriter, r *http.Request) {})
	gQWE.AddRoute(http.MethodGet, "/22", func(w http.ResponseWriter, r *http.Request) {})
	gQWE.AddRoute(http.MethodPatch, "/22", func(w http.ResponseWriter, r *http.Request) {})

	gQWEASD.AddRoute(http.MethodPost, "/11", func(w http.ResponseWriter, r *http.Request) {})
	gQWEASD.AddRoute(http.MethodGet, "/11", func(w http.ResponseWriter, r *http.Request) {})
	gQWEASD.AddRoute(http.MethodPatch, "/11", func(w http.ResponseWriter, r *http.Request) {})
	gQWEASD.AddRoute(http.MethodPost, "/22", func(w http.ResponseWriter, r *http.Request) {})
	gQWEASD.AddRoute(http.MethodGet, "/22", func(w http.ResponseWriter, r *http.Request) {})
	gQWEASD.AddRoute(http.MethodPatch, "/22", func(w http.ResponseWriter, r *http.Request) {})

	err := s.Run()

	if err != nil {
		slog.Error("Running error: ", err)
	}
}
