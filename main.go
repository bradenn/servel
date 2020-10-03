package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"servel/Router"
	"time"
)


func router() {

}

func main() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.HandleFunc("/", Router.ProxyRoute)
	r.HandleFunc("/{path}", Router.ProxyRoute)


	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Failed to listen", err)
	}

}




