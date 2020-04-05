package main

import (
	"log"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/nu50218/nuinfo-syllabus/server"
)

func main() {
	c := server.Config{}
	if err := env.Parse(&c); err != nil {
		log.Fatal(err)
	}
	s := server.New(c)
	log.Fatal(http.ListenAndServe(":8080", s))
}
