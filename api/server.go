package api

import (
	"log"
	"net/http"

	"github.com/Mardiniii/serapis/api/common"
	"github.com/Mardiniii/serapis/api/middlewares"
	"github.com/Mardiniii/serapis/api/routes"
	"github.com/urfave/negroni"
)

// Init inits API server
func Init() {
	var router = routes.Router()

	n := negroni.New()
	n.Use(negroni.HandlerFunc(middlewares.Logger))
	n.Use(negroni.HandlerFunc(middlewares.AuthHeaderValidator))
	n.Use(negroni.HandlerFunc(middlewares.Authenticator))
	n.UseHandler(router)

	println("Creating seed data")
	common.RepoSeedData()

	println("Starting server on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", n))
}